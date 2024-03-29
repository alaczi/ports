module port_domain_service

go 1.21

require (
	github.com/alaczi/ports/logger v0.3.0
	github.com/alaczi/ports/ports v0.3.0
	github.com/alaczi/ports/repository v0.3.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/onsi/ginkgo/v2 v2.13.2
	github.com/onsi/gomega v1.29.0
	go.uber.org/dig v1.17.1
	go.uber.org/mock v0.4.0
	google.golang.org/grpc v1.60.1
)

require (
	github.com/go-logr/logr v1.3.0 // indirect
	github.com/go-task/slim-sprig v0.0.0-20230315185526-52ccab3ef572 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/pprof v0.0.0-20210407192527-94a9f03dee38 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sys v0.14.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	golang.org/x/tools v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231002182017-d307bd883b97 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/alaczi/ports/ports v0.3.0 => ./../../pkg/ports

replace github.com/alaczi/ports/repository v0.3.0 => ./../../pkg/repository

replace github.com/alaczi/ports/logger v0.3.0 => ./../../pkg/logger
