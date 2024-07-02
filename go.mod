module github.com/cloudprober/cloudprober

go 1.21

require (
	bitbucket.org/ubyon/golibs/grpcclient/conn v0.0.0-20240613123936-105e363ca531
	bitbucket.org/ubyon/golibs/logger v0.0.0-20240613123936-105e363ca531
	bitbucket.org/ubyon/golibs/logger/logrus v0.0.0-20240613123936-105e363ca531
	bitbucket.org/ubyon/golibs/ubyon/tssclient v0.0.0-20240613123936-105e363ca531
	cloud.google.com/go/bigquery v1.61.0
	cloud.google.com/go/compute/metadata v0.3.0
	cloud.google.com/go/logging v1.10.0
	cloud.google.com/go/pubsub v1.39.0
	github.com/Masterminds/sprig/v3 v3.2.3
	github.com/aws/aws-sdk-go-v2 v1.30.0
	github.com/aws/aws-sdk-go-v2/config v1.27.21
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.16.8
	github.com/aws/aws-sdk-go-v2/service/cloudwatch v1.39.1
	github.com/fullstorydev/grpcurl v1.9.1
	github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510
	github.com/google/uuid v1.6.0
	github.com/hoisie/redis v0.0.0-20160730154456-b5c6e81454e0
	github.com/jhump/protoreflect v1.16.0
	github.com/kylelemons/godebug v1.1.0
	github.com/lib/pq v1.10.9
	github.com/miekg/dns v1.1.61
	go.opentelemetry.io/otel v1.27.0
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc v1.27.0
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp v1.27.0
	go.opentelemetry.io/otel/exporters/stdout/stdoutmetric v1.27.0
	go.opentelemetry.io/otel/sdk v1.27.0
	go.opentelemetry.io/otel/sdk/metric v1.27.0
	golang.org/x/net v0.26.0
	golang.org/x/oauth2 v0.21.0
	golang.org/x/sys v0.21.0
	google.golang.org/api v0.185.0
	google.golang.org/genproto v0.0.0-20240624140628-dc46fd24d27d
	google.golang.org/grpc v1.64.0
	google.golang.org/protobuf v1.34.2
	sigs.k8s.io/yaml v1.4.0
)

require (
	bitbucket.org/ubyon/core-common/proto/telemetrysrvr/golang v0.0.0-20240603054411-5316caa52ea5 // indirect
	cel.dev/expr v0.15.0 // indirect
	cloud.google.com/go v0.115.0 // indirect
	cloud.google.com/go/auth v0.5.2 // indirect
	cloud.google.com/go/auth/oauth2adapt v0.2.2 // indirect
	cloud.google.com/go/iam v1.1.8 // indirect
	cloud.google.com/go/longrunning v0.5.7 // indirect
	github.com/Masterminds/goutils v1.1.1 // indirect
	github.com/Masterminds/semver/v3 v3.2.1 // indirect
	github.com/apache/arrow/go/v15 v15.0.2 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.17.21 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.3.12 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.6.12 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.8.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.11.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.11.14 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.21.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.25.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.29.1 // indirect
	github.com/aws/smithy-go v1.20.2 // indirect
	github.com/bufbuild/protocompile v0.14.0 // indirect
	github.com/cenkalti/backoff/v4 v4.3.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/cncf/xds/go v0.0.0-20240423153145-555b57ec207b // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/goccy/go-json v0.10.3 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/flatbuffers v24.3.25+incompatible // indirect
	github.com/google/s2a-go v0.1.7 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.3.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.20.0 // indirect
	github.com/huandu/xstrings v1.5.0 // indirect
	github.com/imdario/mergo v0.3.16 // indirect
	github.com/itchyny/timefmt-go v0.1.6 // indirect
	github.com/klauspost/compress v1.17.9 // indirect
	github.com/klauspost/cpuid/v2 v2.2.8 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/pierrec/lz4/v4 v4.1.21 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/shopspring/decimal v1.4.0 // indirect
	github.com/sirupsen/logrus v1.9.0 // indirect
	github.com/spf13/cast v1.6.0 // indirect
	github.com/zeebo/xxh3 v1.0.2 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.52.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.52.0 // indirect
	go.opentelemetry.io/otel/metric v1.27.0 // indirect
	go.opentelemetry.io/otel/trace v1.27.0 // indirect
	go.opentelemetry.io/proto/otlp v1.3.1 // indirect
	golang.org/x/exp v0.0.0-20240613232115-7f521ea00fb8 // indirect
	golang.org/x/time v0.5.0 // indirect
	golang.org/x/xerrors v0.0.0-20231012003039-104605ab7028 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240624140628-dc46fd24d27d // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240624140628-dc46fd24d27d // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

require (
	github.com/census-instrumentation/opencensus-proto v0.4.1 // indirect
	github.com/envoyproxy/go-control-plane v0.12.0 // indirect
	github.com/envoyproxy/protoc-gen-validate v1.0.4 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/googleapis/gax-go/v2 v2.12.5 // indirect
	github.com/itchyny/gojq v0.12.16
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/stretchr/testify v1.9.0
	go.opencensus.io v0.24.0 // indirect
	golang.org/x/crypto v0.24.0 // indirect
	golang.org/x/mod v0.18.0 // indirect
	golang.org/x/sync v0.7.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	golang.org/x/tools v0.22.0
)

replace github.com/cloudprober/cloudprober => github.com/ubyon/cloudprober1 v0.0.0-20240621211417-c227ead67c2c
