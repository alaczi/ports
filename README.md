# Ports

This is a half-baked coding test solution

### How to run
The provided makefile provides multiple utilities to build and execute the code

###### (Re)build the application
```
    make docker-build
```

###### Start the application
```
    make docker-start
```

By default the client listens on the port 8080 while the GRPC endpoint exposed on 50051. This can be changed using environment variables.

###### Stop the application
```
    make docker-stop
```

### What is missing
- Tests are missing completely
- Would have been nice to generalize the GRPC streaming upsert and introduce it to the repository interface
- The docker build is good for the local development, but not for building production code (yet)
- Due lack of time the DTO for Client REST api is the same as the entity model for the repository - really should be different
- More development needed to handle the exit signals while the initial data streaming is still in progress
- Linters, spell checkers