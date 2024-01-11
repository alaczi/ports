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

Install golang using the [official documentation](https://go.dev/doc/install)
Install protoc using the [guide](https://grpc.io/docs/protoc-installation/)
Install golang plugins described [here](https://grpc.io/docs/languages/go/quickstart/)
install [ginkgo]https://onsi.github.io/ginkgo/ for test 

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


###### (Re)build the application
```
    make docker-build
```

###### Start the application
```
    make docker-start
```

By default, the client listens on the port 8080 while the GRPC endpoint exposed on 50051. This can be changed using environment variables.

###### Stop the application
```
    make docker-stop
```

### Todo
- Increase test coverage
- The docker build is good for the local development, but not for building production code (yet)
- Linters