module github.com/observiq/bindplane-agent/exporter/chronicleexporter

go 1.20

require (
	github.com/jarcoal/httpmock v1.3.1
	github.com/observiq/bindplane-agent/expr v1.45.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl v0.94.0
	github.com/stretchr/testify v1.8.4
	go.opentelemetry.io/collector/component v0.94.1
	go.opentelemetry.io/collector/config/configretry v0.94.1
	go.opentelemetry.io/collector/consumer v0.94.1
	go.opentelemetry.io/collector/exporter v0.94.1
	go.opentelemetry.io/collector/pdata v1.1.0
	go.opentelemetry.io/otel/metric v1.23.0
	go.opentelemetry.io/otel/trace v1.23.0
	go.uber.org/zap v1.26.0
	golang.org/x/oauth2 v0.16.0
)

require (
	cloud.google.com/go v0.26.0 // indirect
	github.com/alecthomas/participle/v2 v2.1.1 // indirect
	github.com/antonmedv/expr v1.15.5 // indirect
	github.com/cenkalti/backoff/v4 v4.2.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-viper/mapstructure/v2 v2.0.0-alpha.1 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/iancoleman/strcase v0.3.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/knadh/koanf/providers/confmap v0.1.0 // indirect
	github.com/knadh/koanf/v2 v2.0.2 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/mapstructure v1.5.1-0.20231216201459-8508981c8b6c // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal v0.94.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/objx v0.5.0 // indirect
	go.opentelemetry.io/collector v0.94.1 // indirect
	go.opentelemetry.io/collector/config/configtelemetry v0.94.1 // indirect
	go.opentelemetry.io/collector/confmap v0.94.1 // indirect
	go.opentelemetry.io/collector/extension v0.94.1 // indirect
	go.opentelemetry.io/otel v1.23.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/exp v0.0.0-20240103183307-be819d1f06fc // indirect
	golang.org/x/net v0.20.0 // indirect
	golang.org/x/sys v0.16.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240102182953-50ed04b92917 // indirect
	google.golang.org/grpc v1.61.0 // indirect
	google.golang.org/protobuf v1.32.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/observiq/bindplane-agent/expr => ../../expr
