package grpcserver

import (
	"context"
	"errors"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/samber/lo"
)

var DefaultRecoveryOptions = []grpc_recovery.Option{
	grpc_recovery.WithRecoveryHandler(func(p interface{}) (err error) {
		return status.Errorf(codes.Unknown, "panic triggered: %v", p)
	}),
}

var DefaultCtxTagsOptions = []grpc_ctxtags.Option{
	grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.TagBasedRequestFieldExtractor("log_fields")),
}

var DefaultServerOptions = []grpc.ServerOption{
	grpc.ChainStreamInterceptor(
		grpc_recovery.StreamServerInterceptor(DefaultRecoveryOptions...),
		grpc_ctxtags.StreamServerInterceptor(DefaultCtxTagsOptions...),
		grpc_prometheus.StreamServerInterceptor,
		otelgrpc.StreamServerInterceptor(),
		grpc_auth.StreamServerInterceptor(GrpcSignCheckFunc),
	),
	grpc.ChainUnaryInterceptor(
		grpc_recovery.UnaryServerInterceptor(DefaultRecoveryOptions...),
		grpc_ctxtags.UnaryServerInterceptor(DefaultCtxTagsOptions...),
		grpc_prometheus.UnaryServerInterceptor,
		otelgrpc.UnaryServerInterceptor(),
		grpc_auth.UnaryServerInterceptor(GrpcSignCheckFunc),
	),
}

var SignKey = ""

func GrpcSignCheckFunc(ctx context.Context) (context.Context, error) {
	sign := lo.Must(metadata.FromIncomingContext(ctx)).Get("sign")
	if sign == nil || sign[0] != SignKey {
		return ctx, errors.New("sign checked error")
	}
	return ctx, nil
}
