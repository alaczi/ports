# Ports

A coding test application. 
Two services sharing components.

Domain backend:
A GRPC / Protobuff backend to provide a repository service to store / retrieve port data

Client API
A REST API to expose port data utilizing the domain backend

On startup the client application reads the port data from the provided json file and adds the items to the backend service 

### Setup

Some of the commands assumes that a golang environment was set up on the developer's machine.
Requires golang 1.21.1 version.

- Install golang using the [official documentation](https://go.dev/doc/install)
- Install protoc using the [guide](https://grpc.io/docs/protoc-installation/)
- Install golang plugins described [here](https://grpc.io/docs/languages/go/quickstart/)
- Install [ginkgo](https://onsi.github.io/ginkgo/) for test
- Install [staticcheck](https://github.com/dominikh/go-tools#installation) `go install honnef.co/go/tools/cmd/staticcheck@2023.1.6`

### Application configuration
The services used environment variables for configuration as described in [12 factor app](https://12factor.net/config) 

#### Domain app

The domain app configuration parameters use the "DOMAIN_" prefix. The default values in the list

- DOMAIN_PORT="50051" //defines the port for the GRPC service

#### Client app

The client app configuration parameters use the "CLIENT_" prefix

- CLIENT_SERVERPORT="8080" //the port where the http service for REST endpoints listen
- CLIENT_PORTSERVICEADDR="localhost:50051"  //the address of the grpc backend with port
- CLIENT_DATAFILE="./data/ports.json" //the path to the json file with the initial data

### How to run

The makefile multiple utilities to build, execute, test, lint the code.

###### Run (&build) locally
```shell
    make run
```

###### Run tests
```shell
    make test
```

###### Run linter
```shell
    make lint
```

###### (Re)build the application with docker
```shell
    make docker-build
```

###### Start the application in docker
```shell
    make docker-start
```

###### Stop the docker containers
```shell
    make docker-stop
```

### Todo
- Increase test coverage
- The docker build is good for the local development, but not for building production code (yet)
- Linters