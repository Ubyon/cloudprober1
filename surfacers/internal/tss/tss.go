/*
Package tss implements "tss" surfacer. This surfacer type is in
experimental phase right now.

To use this surfacer, add a stanza similar to the following to your
cloudprober config:

	surfacer {
	  type: tss
		tss_surfacer {
		  connection_string: ""
	  }
	}
*/
package tss

import (
	"context"
    "errors"
    "io/ioutil"
    "os"
	"strconv"
    "sync/atomic"

    "google.golang.org/protobuf/proto"
    "google.golang.org/protobuf/types/known/timestamppb"

	"github.com/cloudprober/cloudprober/logger"
	"github.com/cloudprober/cloudprober/metrics"
	"github.com/cloudprober/cloudprober/surfacers/internal/common/options"
    tsspb "github.com/cloudprober/cloudprober/surfacers/internal/tss/proto"

    gconn "bitbucket.org/ubyon/golibs/grpcclient/conn"
    glgr "bitbucket.org/ubyon/golibs/logger"
    tconn "bitbucket.org/ubyon/golibs/ubyon/tssclient"
	ulrs "bitbucket.org/ubyon/golibs/logger/logrus"
)

// Surfacer structures for writing to tss.
type Surfacer struct {
	cfg             *tsspb.SurfacerConf
	opts            *options.Options
	tssStreamChan   chan *metrics.EventMetrics
    tssClient       *tconn.GrpcTelemetryClientT
    tssClientSeqNo  uint64
	logger          *logger.Logger
}

/*
* Private functions
*/
func (s *Surfacer) getType(em *metrics.EventMetrics)  tsspb.ResourceMetric_Type {
    var pbType tsspb.ResourceMetric_Type
    switch em.Label("ptype") {
    case "http":
        pbType = tsspb.ResourceMetric_HTTP
    default:
        pbType = tsspb.ResourceMetric_INVALID
    }

    return pbType
}

func (s *Surfacer) getKind(em *metrics.EventMetrics)  tsspb.ResourceMetric_Kind {
    var pbKind tsspb.ResourceMetric_Kind
    switch em.Kind {
    case metrics.CUMULATIVE:
        pbKind = tsspb.ResourceMetric_COUNTER
    case metrics.GAUGE:
        pbKind = tsspb.ResourceMetric_GUAGE
    default:
        pbKind = tsspb.ResourceMetric_UNDEFINED
    }

    return pbKind
}

func (s *Surfacer) getLabels(labels map[string]string) []*tsspb.MetricLabel {
    pbLabels := []*tsspb.MetricLabel{}

    for k, v := range labels {
        pbLabel := &tsspb.MetricLabel{
            Key : k,
            Value: v,
            Scope: tsspb.MetricLabel_RESOURCE,
        }

        pbLabels = append(pbLabels, pbLabel)
    }

    return pbLabels
}

func (s *Surfacer) newResourceMetric(em *metrics.EventMetrics, 
                        name, value string, labels map[string]string) *tsspb.ResourceMetric {
    return &tsspb.ResourceMetric{
        Type: s.getType(em),
        Kind: s.getKind(em),
        Name: name,
        Value: value,
        Labels: s.getLabels(labels),
        OrgId: em.Label("org_id"),
        ResourceId: em.Label("resource_id"),
        ResourceName: em.Label("resource_name"),
        GeneratedBy: tsspb.ResourceMetric_TG,
        Timestamp: timestamppb.New(em.Timestamp),
	}
}

func updateLabelMap(labels map[string]string, extraLabels ...[2]string) map[string]string {
	if len(extraLabels) == 0 {
		return labels
	}

    labelsCopy := make(map[string]string)
	for k, v := range labels {
		labelsCopy[k] = v
	}

    for _, extraLabel := range extraLabels {
		labelsCopy[extraLabel[0]] = extraLabel[1]
	}

    return labelsCopy
}

func (s *Surfacer) distToResourceMetrics(d *metrics.DistributionData, em *metrics.EventMetrics,
                            metricName string, labels map[string]string) []*tsspb.ResourceMetric {
	resMetrics := []*tsspb.ResourceMetric{
		s.newResourceMetric(em, metricName+"_sum", strconv.FormatFloat(d.Sum, 'f', -1, 64), labels),
		s.newResourceMetric(em, metricName+"_count", strconv.FormatInt(d.Count, 10), labels),
	}

	var val int64
	for i := range d.LowerBounds {
		val += d.BucketCounts[i]
		var lb string

		if i == len(d.LowerBounds)-1 {
			lb = "+Inf"
		} else {
			lb = strconv.FormatFloat(d.LowerBounds[i+1], 'f', -1, 64)
		}

        labelsWithBucket := updateLabelMap(labels, [2]string{"le", lb})
        brm := s.newResourceMetric(em, metricName+"_bucket", strconv.FormatInt(val, 10), labelsWithBucket)
		resMetrics = append(resMetrics, brm)
	}

	return resMetrics
}

func (s *Surfacer) intMapToResourceMetrics(m *metrics.Map[int64],
                    em *metrics.EventMetrics, metricName string,
                    baseLabels map[string]string) []*tsspb.ResourceMetric {
	resMetrics := []*tsspb.ResourceMetric{}
	for _, k := range m.Keys() {
		labels := updateLabelMap(baseLabels, [2]string{m.MapName, k})
        irm := s.newResourceMetric(em, metricName, metrics.MapValueToString[int64](m.GetKey(k)), labels)
		resMetrics = append(resMetrics, irm)
	}

	return resMetrics
}

func (s *Surfacer) floatMapToResourceMetrics(m *metrics.Map[float64],
                    em *metrics.EventMetrics, metricName string,
                    baseLabels map[string]string) []*tsspb.ResourceMetric {
	resMetrics := []*tsspb.ResourceMetric{}
	for _, k := range m.Keys() {
		labels := updateLabelMap(baseLabels, [2]string{m.MapName, k})
        irm := s.newResourceMetric(em, metricName, metrics.MapValueToString[float64](m.GetKey(k)), labels)
		resMetrics = append(resMetrics, irm)
	}

	return resMetrics
}

func (s *Surfacer) emToResourceMetrics(em *metrics.EventMetrics) *tsspb.ResourceMetricList {
	baseLabels := make(map[string]string)
	for _, k := range em.LabelsKeys() {
		baseLabels[k] = em.Label(k)
	}

	resMetrics := []*tsspb.ResourceMetric{}
	for _, metricName := range em.MetricsKeys() {
		if !s.opts.AllowMetric(metricName) {
			continue
		}

		val := em.Metric(metricName)

		if mapVal, ok := val.(*metrics.Map[int64]); ok {
			resMetrics = append(resMetrics, s.intMapToResourceMetrics(mapVal, em, metricName, baseLabels)...)
			continue
		}

        if mapVal, ok := val.(*metrics.Map[float64]); ok {
			resMetrics = append(resMetrics, s.floatMapToResourceMetrics(mapVal, em, metricName, baseLabels)...)
			continue
		}

		if distVal, ok := val.(*metrics.Distribution); ok {
			resMetrics = append(resMetrics, s.distToResourceMetrics(distVal.Data(), em, metricName, baseLabels)...)
			continue
		}

		if _, ok := val.(metrics.String); ok {
			labels := updateLabelMap(baseLabels, [2]string{"val", val.String()})
			resMetrics = append(resMetrics, s.newResourceMetric(em, metricName, "1", labels))
			continue
		}

		resMetrics = append(resMetrics, s.newResourceMetric(em, metricName, val.String(), baseLabels))
	}

    rml := &tsspb.ResourceMetricList {
        Metrics: resMetrics,
    }

	return rml
}

func (s *Surfacer) streamMetrics(em *metrics.EventMetrics) error {
    if s.tssClient == nil {
        err := errors.New("Invalid TSSClient")
		s.logger.Errorf("Stream metrics err %v", err)
		return err
	}

    s.logger.Infof("Data before conversion %+v", em)
    rml := s.emToResourceMetrics(em)
    // Convert to proto marshal if needed <-- discuss with SD

    /*
    var data [][]byte
    for _, rm := range rms {
        ldata, err := protojson.MarshalOptions { EmitUnpopulated: true,
                                                UseEnumNumbers:  false,
                                              }.Marshal(rm)
	    if err != nil {
	    	s.logger.Errorf("Proto marshal err %v", err)
            return err
        }

        data = append(data, ldata)
    }*/

    s.logger.Infof("Data to be sent %+v", rml)
    jdata, err := proto.Marshal(rml)
    if err != nil {
        s.logger.Errorf("JSON marshal err %v", err)
        return err
    }

    stream, cancelFn, err := s.tssClient.OpenStream()
	defer s.tssClient.CloseStream(stream, cancelFn)
    if err != nil {
        s.logger.Errorf("Open stream err %v", err)
        return err
    }

    topicName := s.cfg.GetTopicName()
    seqNo := atomic.AddUint64(&s.tssClientSeqNo, 1)
    orgId := em.Label("org_id")
	err = s.tssClient.SendMessage(stream, 
                topicName, "",
		        jdata, seqNo, 
                orgId, "")
	if err != nil {
		s.logger.Errorf("OpenStream send message err %v", err)
		return err
	}

	return nil
}

func New(ctx context.Context, config *tsspb.SurfacerConf, logger *logger.Logger) (*Surfacer, error) {
	s := &Surfacer{
		cfg: config,
		logger: logger,
	}

    err := s.init(ctx)
    if err != nil {
        return nil, err
    }

	return s, nil
}

func (s *Surfacer) init(ctx context.Context) error {
    serverAddr := s.cfg.GetConnectionString()
    isTls := s.cfg.GetIsTls()

    grpcConn := gconn.NewConn(serverAddr, isTls)
    ulrs.MkLogger(glgr.LogLevelDebug, "", nil, nil)
    s.tssClient = tconn.NewGrpcClient(serverAddr, grpcConn)
    s.tssClient.SetSecureMode(isTls)
    if isTls {
        caPem, err := ioutil.ReadFile(os.Getenv("CA_FILE"))
        if err != nil {
            s.logger.Errorf("CA file read error %v", err)
            return err
        }

        s.tssClient.SetCertificateAuthority(caPem)
    }

    err := s.tssClient.Start()
	if err != nil {
        s.logger.Errorf("Tss client start err %v", err)
		return err
	}

	s.tssStreamChan = make(chan *metrics.EventMetrics, s.cfg.GetMetricsBufferSize())
	go func() {
		for {
			select {
			case <-ctx.Done():
				s.logger.Infof("Context canceled, stopping the surfacer write loop")
                s.tssClient.Stop()
                s.tssClient = nil
				return
			case em := <-s.tssStreamChan:
				if em.Kind != metrics.CUMULATIVE && em.Kind != metrics.GAUGE {
					continue
				}

                // Note: we may want to batch calls to streamMetrics, as each call results in
				// a stream of data???
                if err := s.streamMetrics(em); err != nil {
					s.logger.Warningf("Error while writing metrics: %v", err)
				}
			}
		}
	}()

	return nil
}

func (s *Surfacer) Write(ctx context.Context, em *metrics.EventMetrics) {
	select {
	case s.tssStreamChan <- em:
	default:
		s.logger.Errorf("TSS surfacer's stream channel is full, dropping new data.")
	}
}
