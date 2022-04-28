// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: campusapis/gate/v1/manager.proto

/*
Package gatev1 is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package gatev1

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

func request_GateManagerService_PostGateAccessCallback_0(ctx context.Context, marshaler runtime.Marshaler, client GateManagerServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq PostGateAccessRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq.Data); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["service"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "service")
	}

	protoReq.Service, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "service", err)
	}

	msg, err := client.PostGateAccessCallback(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_GateManagerService_PostGateAccessCallback_0(ctx context.Context, marshaler runtime.Marshaler, server GateManagerServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq PostGateAccessRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq.Data); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["service"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "service")
	}

	protoReq.Service, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "service", err)
	}

	msg, err := server.PostGateAccessCallback(ctx, &protoReq)
	return msg, metadata, err

}

func request_GateManagerService_PostStudentGateAccess_0(ctx context.Context, marshaler runtime.Marshaler, client GateManagerServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq PostStudentGateAccessRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.PostStudentGateAccess(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_GateManagerService_PostStudentGateAccess_0(ctx context.Context, marshaler runtime.Marshaler, server GateManagerServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq PostStudentGateAccessRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.PostStudentGateAccess(ctx, &protoReq)
	return msg, metadata, err

}

func request_GateManagerService_PostRegisterGateEvent_0(ctx context.Context, marshaler runtime.Marshaler, client GateManagerServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq PostRegisterGateEventRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.PostRegisterGateEvent(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_GateManagerService_PostRegisterGateEvent_0(ctx context.Context, marshaler runtime.Marshaler, server GateManagerServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq PostRegisterGateEventRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.PostRegisterGateEvent(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterGateManagerServiceHandlerServer registers the http handlers for service GateManagerService to "mux".
// UnaryRPC     :call GateManagerServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterGateManagerServiceHandlerFromEndpoint instead.
func RegisterGateManagerServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server GateManagerServiceServer) error {

	mux.Handle("POST", pattern_GateManagerService_PostGateAccessCallback_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		ctx, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/campusapis.gate.v1.GateManagerService/PostGateAccessCallback", runtime.WithHTTPPathPattern("/gate/v1/callback/{service}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_GateManagerService_PostGateAccessCallback_0(ctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_GateManagerService_PostGateAccessCallback_0(ctx, mux, outboundMarshaler, w, req, response_GateManagerService_PostGateAccessCallback_0{resp}, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_GateManagerService_PostStudentGateAccess_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		ctx, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/campusapis.gate.v1.GateManagerService/PostStudentGateAccess", runtime.WithHTTPPathPattern("/campusapis.gate.v1.GateManagerService/PostStudentGateAccess"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_GateManagerService_PostStudentGateAccess_0(ctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_GateManagerService_PostStudentGateAccess_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_GateManagerService_PostRegisterGateEvent_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		ctx, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/campusapis.gate.v1.GateManagerService/PostRegisterGateEvent", runtime.WithHTTPPathPattern("/campusapis.gate.v1.GateManagerService/PostRegisterGateEvent"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_GateManagerService_PostRegisterGateEvent_0(ctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_GateManagerService_PostRegisterGateEvent_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterGateManagerServiceHandlerFromEndpoint is same as RegisterGateManagerServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterGateManagerServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterGateManagerServiceHandler(ctx, mux, conn)
}

// RegisterGateManagerServiceHandler registers the http handlers for service GateManagerService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterGateManagerServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterGateManagerServiceHandlerClient(ctx, mux, NewGateManagerServiceClient(conn))
}

// RegisterGateManagerServiceHandlerClient registers the http handlers for service GateManagerService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "GateManagerServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "GateManagerServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "GateManagerServiceClient" to call the correct interceptors.
func RegisterGateManagerServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client GateManagerServiceClient) error {

	mux.Handle("POST", pattern_GateManagerService_PostGateAccessCallback_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		ctx, err = runtime.AnnotateContext(ctx, mux, req, "/campusapis.gate.v1.GateManagerService/PostGateAccessCallback", runtime.WithHTTPPathPattern("/gate/v1/callback/{service}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_GateManagerService_PostGateAccessCallback_0(ctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_GateManagerService_PostGateAccessCallback_0(ctx, mux, outboundMarshaler, w, req, response_GateManagerService_PostGateAccessCallback_0{resp}, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_GateManagerService_PostStudentGateAccess_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		ctx, err = runtime.AnnotateContext(ctx, mux, req, "/campusapis.gate.v1.GateManagerService/PostStudentGateAccess", runtime.WithHTTPPathPattern("/campusapis.gate.v1.GateManagerService/PostStudentGateAccess"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_GateManagerService_PostStudentGateAccess_0(ctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_GateManagerService_PostStudentGateAccess_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_GateManagerService_PostRegisterGateEvent_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		ctx, err = runtime.AnnotateContext(ctx, mux, req, "/campusapis.gate.v1.GateManagerService/PostRegisterGateEvent", runtime.WithHTTPPathPattern("/campusapis.gate.v1.GateManagerService/PostRegisterGateEvent"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_GateManagerService_PostRegisterGateEvent_0(ctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_GateManagerService_PostRegisterGateEvent_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

type response_GateManagerService_PostGateAccessCallback_0 struct {
	proto.Message
}

func (m response_GateManagerService_PostGateAccessCallback_0) XXX_ResponseBody() interface{} {
	response := m.Message.(*PostGateAccessResponse)
	return response.Message
}

var (
	pattern_GateManagerService_PostGateAccessCallback_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 1, 0, 4, 1, 5, 3}, []string{"gate", "v1", "callback", "service"}, ""))

	pattern_GateManagerService_PostStudentGateAccess_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"campusapis.gate.v1.GateManagerService", "PostStudentGateAccess"}, ""))

	pattern_GateManagerService_PostRegisterGateEvent_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"campusapis.gate.v1.GateManagerService", "PostRegisterGateEvent"}, ""))
)

var (
	forward_GateManagerService_PostGateAccessCallback_0 = runtime.ForwardResponseMessage

	forward_GateManagerService_PostStudentGateAccess_0 = runtime.ForwardResponseMessage

	forward_GateManagerService_PostRegisterGateEvent_0 = runtime.ForwardResponseMessage
)