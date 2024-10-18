// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: gatewayapis/server/v1/manager.proto

/*
Package serverv1 is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package serverv1

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

func request_ServerManageerService_GetHTTPRouterList_0(ctx context.Context, marshaler runtime.Marshaler, client ServerManageerServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetHTTPRouterListRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.GetHTTPRouterList(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_ServerManageerService_GetHTTPRouterList_0(ctx context.Context, marshaler runtime.Marshaler, server ServerManageerServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetHTTPRouterListRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.GetHTTPRouterList(ctx, &protoReq)
	return msg, metadata, err

}

func request_ServerManageerService_GetHTTPRouter_0(ctx context.Context, marshaler runtime.Marshaler, client ServerManageerServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetHTTPRouterRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.GetHTTPRouter(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_ServerManageerService_GetHTTPRouter_0(ctx context.Context, marshaler runtime.Marshaler, server ServerManageerServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetHTTPRouterRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.GetHTTPRouter(ctx, &protoReq)
	return msg, metadata, err

}

func request_ServerManageerService_PostOrPutHTTPRouter_0(ctx context.Context, marshaler runtime.Marshaler, client ServerManageerServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq PostOrPutHTTPRouterRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.PostOrPutHTTPRouter(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_ServerManageerService_PostOrPutHTTPRouter_0(ctx context.Context, marshaler runtime.Marshaler, server ServerManageerServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq PostOrPutHTTPRouterRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.PostOrPutHTTPRouter(ctx, &protoReq)
	return msg, metadata, err

}

func request_ServerManageerService_DeleteHTTPRouter_0(ctx context.Context, marshaler runtime.Marshaler, client ServerManageerServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq DeleteHTTPRouterRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.DeleteHTTPRouter(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_ServerManageerService_DeleteHTTPRouter_0(ctx context.Context, marshaler runtime.Marshaler, server ServerManageerServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq DeleteHTTPRouterRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.DeleteHTTPRouter(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterServerManageerServiceHandlerServer registers the http handlers for service ServerManageerService to "mux".
// UnaryRPC     :call ServerManageerServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterServerManageerServiceHandlerFromEndpoint instead.
// GRPC interceptors will not work for this type of registration. To use interceptors, you must use the "runtime.WithMiddlewares" option in the "runtime.NewServeMux" call.
func RegisterServerManageerServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server ServerManageerServiceServer) error {

	mux.Handle("POST", pattern_ServerManageerService_GetHTTPRouterList_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/gatewayapis.server.v1.ServerManageerService/GetHTTPRouterList", runtime.WithHTTPPathPattern("/gatewayapis.server.v1.ServerManageerService/GetHTTPRouterList"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_ServerManageerService_GetHTTPRouterList_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ServerManageerService_GetHTTPRouterList_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_ServerManageerService_GetHTTPRouter_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/gatewayapis.server.v1.ServerManageerService/GetHTTPRouter", runtime.WithHTTPPathPattern("/gatewayapis.server.v1.ServerManageerService/GetHTTPRouter"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_ServerManageerService_GetHTTPRouter_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ServerManageerService_GetHTTPRouter_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_ServerManageerService_PostOrPutHTTPRouter_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/gatewayapis.server.v1.ServerManageerService/PostOrPutHTTPRouter", runtime.WithHTTPPathPattern("/gatewayapis.server.v1.ServerManageerService/PostOrPutHTTPRouter"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_ServerManageerService_PostOrPutHTTPRouter_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ServerManageerService_PostOrPutHTTPRouter_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_ServerManageerService_DeleteHTTPRouter_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/gatewayapis.server.v1.ServerManageerService/DeleteHTTPRouter", runtime.WithHTTPPathPattern("/gatewayapis.server.v1.ServerManageerService/DeleteHTTPRouter"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_ServerManageerService_DeleteHTTPRouter_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ServerManageerService_DeleteHTTPRouter_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterServerManageerServiceHandlerFromEndpoint is same as RegisterServerManageerServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterServerManageerServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.NewClient(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterServerManageerServiceHandler(ctx, mux, conn)
}

// RegisterServerManageerServiceHandler registers the http handlers for service ServerManageerService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterServerManageerServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterServerManageerServiceHandlerClient(ctx, mux, NewServerManageerServiceClient(conn))
}

// RegisterServerManageerServiceHandlerClient registers the http handlers for service ServerManageerService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "ServerManageerServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "ServerManageerServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "ServerManageerServiceClient" to call the correct interceptors. This client ignores the HTTP middlewares.
func RegisterServerManageerServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client ServerManageerServiceClient) error {

	mux.Handle("POST", pattern_ServerManageerService_GetHTTPRouterList_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/gatewayapis.server.v1.ServerManageerService/GetHTTPRouterList", runtime.WithHTTPPathPattern("/gatewayapis.server.v1.ServerManageerService/GetHTTPRouterList"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_ServerManageerService_GetHTTPRouterList_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ServerManageerService_GetHTTPRouterList_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_ServerManageerService_GetHTTPRouter_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/gatewayapis.server.v1.ServerManageerService/GetHTTPRouter", runtime.WithHTTPPathPattern("/gatewayapis.server.v1.ServerManageerService/GetHTTPRouter"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_ServerManageerService_GetHTTPRouter_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ServerManageerService_GetHTTPRouter_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_ServerManageerService_PostOrPutHTTPRouter_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/gatewayapis.server.v1.ServerManageerService/PostOrPutHTTPRouter", runtime.WithHTTPPathPattern("/gatewayapis.server.v1.ServerManageerService/PostOrPutHTTPRouter"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_ServerManageerService_PostOrPutHTTPRouter_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ServerManageerService_PostOrPutHTTPRouter_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_ServerManageerService_DeleteHTTPRouter_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/gatewayapis.server.v1.ServerManageerService/DeleteHTTPRouter", runtime.WithHTTPPathPattern("/gatewayapis.server.v1.ServerManageerService/DeleteHTTPRouter"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_ServerManageerService_DeleteHTTPRouter_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ServerManageerService_DeleteHTTPRouter_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_ServerManageerService_GetHTTPRouterList_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"gatewayapis.server.v1.ServerManageerService", "GetHTTPRouterList"}, ""))

	pattern_ServerManageerService_GetHTTPRouter_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"gatewayapis.server.v1.ServerManageerService", "GetHTTPRouter"}, ""))

	pattern_ServerManageerService_PostOrPutHTTPRouter_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"gatewayapis.server.v1.ServerManageerService", "PostOrPutHTTPRouter"}, ""))

	pattern_ServerManageerService_DeleteHTTPRouter_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"gatewayapis.server.v1.ServerManageerService", "DeleteHTTPRouter"}, ""))
)

var (
	forward_ServerManageerService_GetHTTPRouterList_0 = runtime.ForwardResponseMessage

	forward_ServerManageerService_GetHTTPRouter_0 = runtime.ForwardResponseMessage

	forward_ServerManageerService_PostOrPutHTTPRouter_0 = runtime.ForwardResponseMessage

	forward_ServerManageerService_DeleteHTTPRouter_0 = runtime.ForwardResponseMessage
)
