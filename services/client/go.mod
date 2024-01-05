module client

go 1.21

require (
	github.com/alaczi/ports/ports v0.1.0
	github.com/alaczi/ports/repository v0.1.0
	google.golang.org/grpc v1.60.1
	github.com/gorilla/mux v1.8.1
	github.com/kelseyhightower/envconfig v1.4.0
)

replace github.com/alaczi/ports/ports v0.1.0 => ./../../pkg/ports
replace github.com/alaczi/ports/repository v0.1.0 => ./../../pkg/repository