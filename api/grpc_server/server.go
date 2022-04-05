package grpc_server

import (
	"log"
	"net/http"
	"reflect"

	"google.golang.org/grpc"
)

// Backend abstracts a registerable GRPC
type Backend interface {
	RegisterGRPC(*grpc.Server)
}

var server *grpc.Server

func GetInstance() *grpc.Server {
	return server
}

// NewService initializes a new grpc service
func NewService(opts []grpc.ServerOption, backends ...Backend) http.HandlerFunc {
	instance := grpc.NewServer(opts...)

	// Register all grpc backends
	for _, b := range backends {
		b.RegisterGRPC(instance)
		log.Printf("grpc server %s init successfully", reflect.TypeOf(b))
	}
	log.Println("grpc server is running")
	return instance.ServeHTTP
}
