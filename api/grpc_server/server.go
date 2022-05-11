package grpcserver

import (
	"log"
	"reflect"

	"google.golang.org/grpc"
)

// Backend abstracts a registerable GRPC
type Backend interface {
	RegisterGRPC(*grpc.Server)
}

var instance *grpc.Server

func GetInstance() *grpc.Server {
	return instance
}

// NewService initializes a new grpc service
func NewService(backends ...Backend) *grpc.Server {
	instance = grpc.NewServer(DefaultServerOptions...)

	// Register all grpc backends
	for _, b := range backends {
		b.RegisterGRPC(instance)
		log.Printf("grpc server %s init successfully", reflect.TypeOf(b))
	}
	log.Println("grpc server is running")
	return instance
}
