module github.com/observiq/bindplane-otel-collector/internal/rehydration

go 1.23.6

require (
	github.com/observiq/bindplane-otel-collector/internal/testutils v1.71.4
	github.com/stretchr/testify v1.10.0
	go.opentelemetry.io/collector/component v0.120.0
	go.opentelemetry.io/collector/consumer v1.26.0
	go.opentelemetry.io/collector/consumer/consumertest v0.120.0
	go.opentelemetry.io/collector/extension/xextension v0.120.0
	go.opentelemetry.io/collector/pdata v1.26.0
)

require go.opentelemetry.io/collector/consumer/xconsumer v0.120.0 // indirect

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/objx v0.5.2 // indirect
	go.opentelemetry.io/collector/extension v0.120.0 // indirect
	go.opentelemetry.io/collector/pdata/pprofile v0.120.0 // indirect
	go.opentelemetry.io/collector/pipeline v0.120.0
	go.opentelemetry.io/otel v1.34.0 // indirect
	go.opentelemetry.io/otel/metric v1.34.0 // indirect
	go.opentelemetry.io/otel/trace v1.34.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
	golang.org/x/net v0.33.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241202173237-19429a94021a // indirect
	google.golang.org/grpc v1.70.0 // indirect
	google.golang.org/protobuf v1.36.5 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/observiq/bindplane-otel-collector/internal/testutils => ../testutils
