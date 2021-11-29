
all: generate-proto

docker-run:
	docker run -v $(shell pwd):/go/src/github.com/hduhelp/api_open_sdk -w /go/src/github.com/hduhelp/api_open_sdk --rm proto-builder sh -c 'make generate-proto'

generate-proto:
	protoc --proto_path=$(GOPATH)/src:. --go_out=$(GOPATH)/src --go-grpc_out=$(GOPATH)/src baseStaff/baseStaff.proto
	protoc --proto_path=$(GOPATH)/src:. --go_out=$(GOPATH)/src --go-grpc_out=$(GOPATH)/src baseStaff/student/student.proto
	protoc --proto_path=$(GOPATH)/src:. --go_out=$(GOPATH)/src --go-grpc_out=$(GOPATH)/src baseStaff/person/person.proto
	protoc --proto_path=$(GOPATH)/src:. --go_out=$(GOPATH)/src --go-grpc_out=$(GOPATH)/src baseStaff/teacher/teacher.proto
	protoc --proto_path=$(GOPATH)/src:. --go_out=$(GOPATH)/src --go-grpc_out=$(GOPATH)/src teaching/teaching.proto
	protoc --proto_path=$(GOPATH)/src:. --go_out=$(GOPATH)/src --go-grpc_out=$(GOPATH)/src school/school.proto
	protoc --proto_path=$(GOPATH)/src:. --go_out=$(GOPATH)/src --go-grpc_out=$(GOPATH)/src schoolTime/schoolTime.proto
