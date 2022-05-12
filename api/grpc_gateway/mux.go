package grpcgateway

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"reflect"

	"github.com/felixge/httpsnoop"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
)

var DefaultServeMuxOption = []runtime.ServeMuxOption{
	runtime.WithErrorHandler(DefaultErrorHandler),
	runtime.WithRoutingErrorHandler(DefaultRoutingErrorHandler),
	runtime.WithIncomingHeaderMatcher(DefaultHeaderWarp),
	runtime.WithMarshalerOption("application/jsonpb", &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			EmitUnpopulated: true,
			UseProtoNames:   true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: false,
		},
	}),
}

func NewMixedPortServeMux(ctx context.Context, endpoint string, registers ...Register) (*runtime.ServeMux, error) {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	grpcEndpoint, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return nil, fmt.Errorf("dial to grpc server %s error: %s", endpoint, err)
	}

	mux := runtime.NewServeMux(DefaultServeMuxOption...)

	for _, r := range registers {
		if err := r(ctx, mux, grpcEndpoint); err != nil {
			return nil, fmt.Errorf("grpc gateway %s init error: %s", reflect.TypeOf(r), err)
		}
	}
	return mux, nil
}

func GRPCGatewayListener(mux cmux.CMux) net.Listener {
	return mux.Match(cmux.HTTP1Fast())
}

func GRPCServerListener(mux cmux.CMux) net.Listener {
	return mux.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
}

func DefaultHTTPServer(handler http.Handler) *http.Server {
	return HTTPServer(handler, WithLogger, WithResponseWriter)
}

func HTTPServer(handler http.Handler, warps ...func(http.Handler) http.Handler) *http.Server {
	for _, warp := range warps {
		handler = warp(handler)
	}
	return &http.Server{Handler: handler}
}

func WithLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		m := httpsnoop.CaptureMetrics(handler, writer, request)
		log.Printf("http[%d]-- %s -- %s\n", m.Code, m.Duration, request.URL.Path)
	})
}
