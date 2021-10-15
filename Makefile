
all: generate-micro

generate-micro:
	protoc --proto_path=$(GOPATH)/src:. --go_out=$(GOPATH)/src --go-grpc_out=$(GOPATH)/src github.com/hduhelp/api_open_sdk/schedule/schedule.proto
	protoc --proto_path=$(GOPATH)/src:. --go_out=$(GOPATH)/src --go-grpc_out=$(GOPATH)/src github.com/hduhelp/api_open_sdk/baseStaff/baseStaff.proto
	protoc --proto_path=$(GOPATH)/src:. --go_out=$(GOPATH)/src --go-grpc_out=$(GOPATH)/src github.com/hduhelp/api_open_sdk/schoolTime/schoolTime.proto
