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
	"database/sql"
	"encoding/json"
    "io/ioutil"
    "os"
	"strconv"
	"time"

	"github.com/cloudprober/cloudprober/logger"
	"github.com/cloudprober/cloudprober/metrics"
	"github.com/cloudprober/cloudprober/surfacers/internal/common/options"
    tsspb "github.com/cloudprober/cloudprober/surfacers/internal/tss/proto"
	"github.com/lib/pq"

    "bitbucket.org/ubyon/golibs/ubyon/tssclient"
)

// Surfacer structures for writing to tss.
type Surfacer struct {
	config        *tsspb.SurfacerConf
	opts          *options.Options
	tssStreamChan chan *metrics.EventMetrics
    tssClient     *tssclient.GrpcTelemetryClientT
	logger        *logger.Logger
}

type Metric struct {
    ResourceMetrics []*tsspb.ResourceMetric
}

/*
* Private functions
*/
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

func (s *Surfacer) getKind(em *metrics.EventMetrics)  pb.ResourceMetric_Kind {
    switch em.Kind {
    case metrics.CUMULATIVE:
        return pb.ResourceMetric_COUNTER
    case metrics.GAUGE:
        return pb.ResourceMetric_GUAGE
    default:
        return pb.ResourceMetric_UNKNOWN
    }
}


func (s *Surfacer) newResourceMetric(em *metrics.EventMetrics, 
                        name, value string, labels map[string]string) *tsspb.ResourceMetric {
    //TODO SM (get org/resource info and add it to resource metric)
    var pbType tsspb.ResourceMetric_Type
    switch rm.Type {
    case HttpResourceMetricType:
        pbType = pb.ResourceMetric_HTTP
    case HttpsResourceMetricType:
        pbType = pb.ResourceMetric_HTTPS
    default:
        pbType = pb.ResourceMetric_INVALID
    }

    return &tsspb.ResourceMetric{
        Type: s.getType(em),
        Kind: s.getKind(em)
	}
}

func (s *Surfacer) distToResourceMetrics(d *metrics.DistributionData, em *metrics.EventMetrics,
                            metricName string, labels map[string]string) []*tsspb.ResourceMetric {
	resMerics := []*tsspb.ResourceMetric{
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
		resMetrics = append(resMerics, brm)
	}

	return resMetrics
}

func (s *Surfacer) mapToResourceMetrics[T int64 | float64](m *metrics.Map[T],
                    em *metrics.EventMetrics, metricName string,
                    baseLabels map[string]string) []*tsspb.ResourceMetric {
	resMetrics := []*tsspb.ResourceMetric{}
	for _, k := range m.Keys() {
		labels := updateLabelMap(baseLabels, [2]string{m.MapName, k})
        irm := s.newResourceMetric(em, metricName, metrics.MapValueToString[T](m.GetKey(k)), labels)
		resMetrics = append(resMetrics, irm)
	}

	return resMetrics
}

func (s *Surfacer) emToResourceMetrics(em *metrics.EventMetrics) []resourceMetric {
	baseLabels := make(map[string]string)
	for _, k := range em.LabelsKeys() {
		baseLabels[k] = em.Label(k)
	}

	resMetrics := []resourceMetric{}
	for _, metricName := range em.MetricsKeys() {
		if !s.opts.AllowMetric(metricName) {
			continue
		}

		val := em.Metric(metricName)

		if mapVal, ok := val.(*metrics.Map[int64]); ok {
			resMetrics = append(resMetrics, s.mapToResourceMetrics(mapVal, em, metricName, baseLabels)...)
			continue
		}

        if mapVal, ok := val.(*metrics.Map[float64]); ok {
			resMetrics = append(resMetrics, s.mapToResourceMetrics(mapVal, em, metricName, baseLabels)...)
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
	return resMetrics
}

func (s *Surfacer) streamMetrics(em *metrics.EventMetrics) error {
    //TODO stream the data
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
    caPem, err := ioutil.ReadFile(os.Getenv("CA_FILE"))
    if err != nil {
        s.logger.Fatalf("Failed to read ca file: %v", err)
    }

    s.tssClient := tssclient.NewGrpcClient(serverAddr, gConn)
    s.tssClient.SetSecureMode(isTls)
    s.tssCient.SetCertificateAuthority(caPem)
    err := s.tssClient.Start()
	if err != nil {
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
