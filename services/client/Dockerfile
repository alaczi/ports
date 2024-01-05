FROM golang:1.21.5 as build-env

RUN apt-get update && apt-get install -y --no-install-recommends protobuf-compiler
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

WORKDIR /app

ADD . /app
WORKDIR /app

ARG CGO_ENABLED=0

RUN make build-client

FROM gcr.io/distroless/static-debian12
COPY --from=build-env /app/build/client /
COPY ./data /data

CMD ["/client"]