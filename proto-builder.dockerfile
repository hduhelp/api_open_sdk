FROM golang:latest

RUN apt-get update
RUN apt-get -y install git unzip build-essential autoconf libtool make protobuf-compiler

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io,direct

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

ENV PATH="$PATH:$GOPATH/bin"
