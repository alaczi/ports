.PHONY: build

generate_pb:
	protoc --experimental_allow_proto3_optional --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pkg/ports/ports.proto

clean:
	rm -f ./build/*

build-client: clean
	go build -C ./services/client -o ../../build/ -ldflags "-s -w"

build-domain: clean
	go build -C ./services/domain -o ../../build/ -ldflags "-s -w"

build: build-client build-domain
	go build -C ./services/domain -o ../../build/
	go build -C ./services/client -o ../../build/

run: build
	./build/port_domain_service &
	GOMEMLIMIT=200MiB ./build/client &

stop:
	killall port_domain_service
	killall client

fmt:
	go fmt -C ./pkg/ports
	go mod tidy -C ./pkg/ports
	go fmt -C ./pkg/repository
	go mod tidy -C ./pkg/repository
	go fmt -C ./services/client
	go mod tidy -C ./services/client
	go fmt -C ./services/domain
	go mod tidy -C ./services/domain

docker-build-client: build-client
	docker build ./ -f ./services/client/Dockerfile -t client

docker-build-domain: build-domain
	docker build ./ -f ./services/domain/Dockerfile -t domain

docker-rebuild:
	docker-compose build

docker-start:
	docker-compose up --detach

docker-stop:
	docker-compose stop