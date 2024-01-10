module port_domain_service

go 1.21

require (
	github.com/alaczi/ports/ports v0.2.0
	github.com/alaczi/ports/repository v0.2.0
	github.com/kelseyhightower/envconfig v1.4.0
	go.uber.org/dig v1.17.1
	google.golang.org/grpc v1.60.1
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.16.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231002182017-d307bd883b97 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)

replace github.com/alaczi/ports/ports v0.2.0 => ./../../pkg/ports

replace github.com/alaczi/ports/repository v0.2.0 => ./../../pkg/repository