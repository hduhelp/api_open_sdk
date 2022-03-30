// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: campusapis/staff/v1/freshman.proto

/*
Package staffv1 is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package staffv1

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

var (
	filter_FreshmanService_GetFreshmanBaseInfo_0 = &utilities.DoubleArray{Encoding: map[string]int{}, Base: []int(nil), Check: []int(nil)}
)

func request_FreshmanService_GetFreshmanBaseInfo_0(ctx context.Context, marshaler runtime.Marshaler, client FreshmanServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq FreshmanKeywordRequest
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_FreshmanService_GetFreshmanBaseInfo_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.GetFreshmanBaseInfo(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_FreshmanService_GetFreshmanBaseInfo_0(ctx context.Context, marshaler runtime.Marshaler, server FreshmanServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq FreshmanKeywordRequest
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_FreshmanService_GetFreshmanBaseInfo_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.GetFreshmanBaseInfo(ctx, &protoReq)
	return msg, metadata, err

}

var (
	filter_FreshmanService_GetFreshmanBaseInfo_1 = &utilities.DoubleArray{Encoding: map[string]int{}, Base: []int(nil), Check: []int(nil)}
)

func request_FreshmanService_GetFreshmanBaseInfo_1(ctx context.Context, marshaler runtime.Marshaler, client FreshmanServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq FreshmanKeywordRequest
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_FreshmanService_GetFreshmanBaseInfo_1); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.GetFreshmanBaseInfo(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_FreshmanService_GetFreshmanBaseInfo_1(ctx context.Context, marshaler runtime.Marshaler, server FreshmanServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq FreshmanKeywordRequest
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_FreshmanService_GetFreshmanBaseInfo_1); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.GetFreshmanBaseInfo(ctx, &protoReq)
	return msg, metadata, err

}

var (
	filter_FreshmanService_GetFreshmanInfo_0 = &utilities.DoubleArray{Encoding: map[string]int{}, Base: []int(nil), Check: []int(nil)}
)

func request_FreshmanService_GetFreshmanInfo_0(ctx context.Context, marshaler runtime.Marshaler, client FreshmanServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq FreshmanKeywordRequest
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_FreshmanService_GetFreshmanInfo_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.GetFreshmanInfo(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_FreshmanService_GetFreshmanInfo_0(ctx context.Context, marshaler runtime.Marshaler, server FreshmanServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq FreshmanKeywordRequest
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_FreshmanService_GetFreshmanInfo_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.GetFreshmanInfo(ctx, &protoReq)
	return msg, metadata, err

}

var (
	filter_FreshmanService_GetFreshmanInfo_1 = &utilities.DoubleArray{Encoding: map[string]int{}, Base: []int(nil), Check: []int(nil)}
)

func request_FreshmanService_GetFreshmanInfo_1(ctx context.Context, marshaler runtime.Marshaler, client FreshmanServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq FreshmanKeywordRequest
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_FreshmanService_GetFreshmanInfo_1); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.GetFreshmanInfo(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_FreshmanService_GetFreshmanInfo_1(ctx context.Context, marshaler runtime.Marshaler, server FreshmanServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq FreshmanKeywordRequest
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_FreshmanService_GetFreshmanInfo_1); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.GetFreshmanInfo(ctx, &protoReq)
	return msg, metadata, err

}

var (
	filter_FreshmanService_GetFreshmanRoommates_0 = &utilities.DoubleArray{Encoding: map[string]int{}, Base: []int(nil), Check: []int(nil)}
)

func request_FreshmanService_GetFreshmanRoommates_0(ctx context.Context, marshaler runtime.Marshaler, client FreshmanServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq FreshmanKeywordRequest
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_FreshmanService_GetFreshmanRoommates_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.GetFreshmanRoommates(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_FreshmanService_GetFreshmanRoommates_0(ctx context.Context, marshaler runtime.Marshaler, server FreshmanServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq FreshmanKeywordRequest
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_FreshmanService_GetFreshmanRoommates_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.GetFreshmanRoommates(ctx, &protoReq)
	return msg, metadata, err

}

var (
	filter_FreshmanService_GetFreshmanRoommates_1 = &utilities.DoubleArray{Encoding: map[string]int{}, Base: []int(nil), Check: []int(nil)}
)

func request_FreshmanService_GetFreshmanRoommates_1(ctx context.Context, marshaler runtime.Marshaler, client FreshmanServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq FreshmanKeywordRequest
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_FreshmanService_GetFreshmanRoommates_1); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.GetFreshmanRoommates(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_FreshmanService_GetFreshmanRoommates_1(ctx context.Context, marshaler runtime.Marshaler, server FreshmanServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq FreshmanKeywordRequest
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_FreshmanService_GetFreshmanRoommates_1); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.GetFreshmanRoommates(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterFreshmanServiceHandlerServer registers the http handlers for service FreshmanService to "mux".
// UnaryRPC     :call FreshmanServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterFreshmanServiceHandlerFromEndpoint instead.
func RegisterFreshmanServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server FreshmanServiceServer) error {

	mux.Handle("GET", pattern_FreshmanService_GetFreshmanBaseInfo_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		ctx, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/campusapis.staff.v1.FreshmanService/GetFreshmanBaseInfo", runtime.WithHTTPPathPattern("/staff/v1/freshman/base"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_FreshmanService_GetFreshmanBaseInfo_0(ctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_FreshmanService_GetFreshmanBaseInfo_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_FreshmanService_GetFreshmanBaseInfo_1, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		ctx, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/campusapis.staff.v1.FreshmanService/GetFreshmanBaseInfo", runtime.WithHTTPPathPattern("/freshman/base"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_FreshmanService_GetFreshmanBaseInfo_1(ctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_FreshmanService_GetFreshmanBaseInfo_1(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_FreshmanService_GetFreshmanInfo_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		ctx, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/campusapis.staff.v1.FreshmanService/GetFreshmanInfo", runtime.WithHTTPPathPattern("/staff/v1/freshman/info"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_FreshmanService_GetFreshmanInfo_0(ctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_FreshmanService_GetFreshmanInfo_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_FreshmanService_GetFreshmanInfo_1, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		ctx, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/campusapis.staff.v1.FreshmanService/GetFreshmanInfo", runtime.WithHTTPPathPattern("/freshman/info"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_FreshmanService_GetFreshmanInfo_1(ctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_FreshmanService_GetFreshmanInfo_1(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_FreshmanService_GetFreshmanRoommates_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		ctx, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/campusapis.staff.v1.FreshmanService/GetFreshmanRoommates", runtime.WithHTTPPathPattern("/staff/v1/freshman/roommate"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_FreshmanService_GetFreshmanRoommates_0(ctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_FreshmanService_GetFreshmanRoommates_0(ctx, mux, outboundMarshaler, w, req, response_FreshmanService_GetFreshmanRoommates_0{resp}, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_FreshmanService_GetFreshmanRoommates_1, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		ctx, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/campusapis.staff.v1.FreshmanService/GetFreshmanRoommates", runtime.WithHTTPPathPattern("/freshman/roommate"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_FreshmanService_GetFreshmanRoommates_1(ctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_FreshmanService_GetFreshmanRoommates_1(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterFreshmanServiceHandlerFromEndpoint is same as RegisterFreshmanServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterFreshmanServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
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

	return RegisterFreshmanServiceHandler(ctx, mux, conn)
}

// RegisterFreshmanServiceHandler registers the http handlers for service FreshmanService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterFreshmanServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterFreshmanServiceHandlerClient(ctx, mux, NewFreshmanServiceClient(conn))
}

// RegisterFreshmanServiceHandlerClient registers the http handlers for service FreshmanService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "FreshmanServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "FreshmanServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "FreshmanServiceClient" to call the correct interceptors.
func RegisterFreshmanServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client FreshmanServiceClient) error {

	mux.Handle("GET", pattern_FreshmanService_GetFreshmanBaseInfo_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		ctx, err = runtime.AnnotateContext(ctx, mux, req, "/campusapis.staff.v1.FreshmanService/GetFreshmanBaseInfo", runtime.WithHTTPPathPattern("/staff/v1/freshman/base"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_FreshmanService_GetFreshmanBaseInfo_0(ctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_FreshmanService_GetFreshmanBaseInfo_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_FreshmanService_GetFreshmanBaseInfo_1, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		ctx, err = runtime.AnnotateContext(ctx, mux, req, "/campusapis.staff.v1.FreshmanService/GetFreshmanBaseInfo", runtime.WithHTTPPathPattern("/freshman/base"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_FreshmanService_GetFreshmanBaseInfo_1(ctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_FreshmanService_GetFreshmanBaseInfo_1(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_FreshmanService_GetFreshmanInfo_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		ctx, err = runtime.AnnotateContext(ctx, mux, req, "/campusapis.staff.v1.FreshmanService/GetFreshmanInfo", runtime.WithHTTPPathPattern("/staff/v1/freshman/info"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_FreshmanService_GetFreshmanInfo_0(ctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_FreshmanService_GetFreshmanInfo_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_FreshmanService_GetFreshmanInfo_1, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		ctx, err = runtime.AnnotateContext(ctx, mux, req, "/campusapis.staff.v1.FreshmanService/GetFreshmanInfo", runtime.WithHTTPPathPattern("/freshman/info"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_FreshmanService_GetFreshmanInfo_1(ctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_FreshmanService_GetFreshmanInfo_1(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_FreshmanService_GetFreshmanRoommates_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		ctx, err = runtime.AnnotateContext(ctx, mux, req, "/campusapis.staff.v1.FreshmanService/GetFreshmanRoommates", runtime.WithHTTPPathPattern("/staff/v1/freshman/roommate"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_FreshmanService_GetFreshmanRoommates_0(ctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_FreshmanService_GetFreshmanRoommates_0(ctx, mux, outboundMarshaler, w, req, response_FreshmanService_GetFreshmanRoommates_0{resp}, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_FreshmanService_GetFreshmanRoommates_1, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		ctx, err = runtime.AnnotateContext(ctx, mux, req, "/campusapis.staff.v1.FreshmanService/GetFreshmanRoommates", runtime.WithHTTPPathPattern("/freshman/roommate"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_FreshmanService_GetFreshmanRoommates_1(ctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_FreshmanService_GetFreshmanRoommates_1(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

type response_FreshmanService_GetFreshmanRoommates_0 struct {
	proto.Message
}

func (m response_FreshmanService_GetFreshmanRoommates_0) XXX_ResponseBody() interface{} {
	response := m.Message.(*GetFreshmanRoommatesResponse)
	return response.Items
}

var (
	pattern_FreshmanService_GetFreshmanBaseInfo_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"staff", "v1", "freshman", "base"}, ""))

	pattern_FreshmanService_GetFreshmanBaseInfo_1 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"freshman", "base"}, ""))

	pattern_FreshmanService_GetFreshmanInfo_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"staff", "v1", "freshman", "info"}, ""))

	pattern_FreshmanService_GetFreshmanInfo_1 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"freshman", "info"}, ""))

	pattern_FreshmanService_GetFreshmanRoommates_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"staff", "v1", "freshman", "roommate"}, ""))

	pattern_FreshmanService_GetFreshmanRoommates_1 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"freshman", "roommate"}, ""))
)

var (
	forward_FreshmanService_GetFreshmanBaseInfo_0 = runtime.ForwardResponseMessage

	forward_FreshmanService_GetFreshmanBaseInfo_1 = runtime.ForwardResponseMessage

	forward_FreshmanService_GetFreshmanInfo_0 = runtime.ForwardResponseMessage

	forward_FreshmanService_GetFreshmanInfo_1 = runtime.ForwardResponseMessage

	forward_FreshmanService_GetFreshmanRoommates_0 = runtime.ForwardResponseMessage

	forward_FreshmanService_GetFreshmanRoommates_1 = runtime.ForwardResponseMessage
)
