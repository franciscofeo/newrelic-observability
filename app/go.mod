module crudapp

go 1.21

toolchain go1.22.2

require (
	github.com/lib/pq v1.10.9
	github.com/newrelic/go-agent/v3/integrations/logcontext-v2/logWriter v1.0.1
)

require github.com/newrelic/go-agent/v3/integrations/logcontext-v2/nrwriter v1.0.0 // indirect

require (
	github.com/newrelic/go-agent/v3 v3.36.0
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	golang.org/x/text v0.15.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240528184218-531527333157 // indirect
	google.golang.org/grpc v1.65.0 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
)
