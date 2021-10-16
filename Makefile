
all: generate-micro

generate-micro:
	protoc --proto_path=$(GOPATH)/src:. --go_out=$(GOPATH)/src --go-grpc_out=$(GOPATH)/src schedule/schedule.proto
	protoc --proto_path=$(GOPATH)/src:. --go_out=$(GOPATH)/src --go-grpc_out=$(GOPATH)/src baseStaff/baseStaff.proto
	protoc --proto_path=$(GOPATH)/src:. --go_out=$(GOPATH)/src --go-grpc_out=$(GOPATH)/src schoolTime/schoolTime.proto
