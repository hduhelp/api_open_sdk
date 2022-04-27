package grpcclient

import (
	"context"
	"crypto/x509"
	"log"

	"github.com/samber/lo"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

var defaultEndpoint = "gapi.hduhelp.com:443"

func Conn(ctx context.Context, endpoints ...string) grpc.ClientConnInterface {
	endpoint := lo.Ternary(len(endpoints) != 0,
		endpoints[0],
		defaultEndpoint,
	)
	certPool, err := x509.SystemCertPool()
	if err != nil {
		log.Fatalf("failed to load credentials: %v", err)
	}
	creds := credentials.NewClientTLSFromCert(certPool, "gapi.hduhelp.com")
	conn, err := grpc.DialContext(ctx, endpoint, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal("grpc client did not connect", zap.Error(err))
	}
	return conn
}

func WithToken(ctx context.Context, token string) context.Context {
	md := metadata.New(map[string]string{"authorization": "token " + token})
	return metadata.NewOutgoingContext(ctx, md)
}
