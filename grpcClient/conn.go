package grpcclient

import (
	"context"
	"crypto/x509"
	"github.com/hduhelp/api_open_sdk/env"
	"google.golang.org/grpc/credentials"
	"log"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var defaultEndpoint = "gateway-service.hduhelp:9101"

func Conn(ctx context.Context, endpoints ...string) (conn *grpc.ClientConn) {
	var endpoint string
	if len(endpoints) != 0 {
		endpoint = endpoints[0]
	} else {
		endpoint = defaultEndpoint
	}
	var err error
	if !env.IsProd() {
		certPool, err := x509.SystemCertPool()
		if err != nil {
			log.Fatalf("failed to load credentials: %v", err)
		}
		creds := credentials.NewClientTLSFromCert(certPool, "127.0.0.1")
		conn, err = grpc.DialContext(ctx, endpoint, grpc.WithTransportCredentials(creds))
	} else {
		conn, err = grpc.DialContext(ctx, endpoint, grpc.WithInsecure())
	}
	if err != nil || conn == nil {
		log.Fatal("grpc client did not connect", zap.Error(err))
	}
	return conn
}

func WithToken(ctx context.Context, token string) context.Context {
	md := metadata.New(map[string]string{"authorization": "token " + token})
	return metadata.NewOutgoingContext(ctx, md)
}
