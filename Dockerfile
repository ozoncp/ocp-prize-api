FROM golang:1.16

RUN apt update

RUN apt install -y git
RUN apt install -y ca-certificates

RUN apt install -y protobuf-compiler

COPY . /home/user/github.com/ozoncp/ocp-prize-api

WORKDIR /home/user/github.com/ozoncp/ocp-prize-api

RUN go get github.com/golangci/golangci-lint/cmd/golangci-lint

RUN go get github.com/Shopify/sarama
RUN go get github.com/prometheus/client_golang/prometheus

RUN make deps

RUN go get github.com/envoyproxy/protoc-gen-validate
RUN go install github.com/envoyproxy/protoc-gen-validate

RUN go get github.com/opentracing/opentracing-go

RUN make lint
RUN make build

EXPOSE 82