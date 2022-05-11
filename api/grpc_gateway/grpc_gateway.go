package grpcgateway

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/textproto"

	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/hduhelp/api_open_sdk/campusapis"
	"github.com/samber/lo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Register func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error

func RoutingErrorHandlerFinal(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, r *http.Request, httpStatus int) {
	sterr := status.Error(codes.Internal, "Unexpected routing error")
	switch httpStatus {
	case http.StatusBadRequest:
		sterr = status.Error(codes.InvalidArgument, "")
	case http.StatusMethodNotAllowed:
		sterr = status.Error(codes.Unimplemented, "")
	case http.StatusNotFound:
		sterr = status.Error(codes.NotFound, "router not found")
	}
	ErrorHandler(ctx, mux, marshaler, w, r, sterr)
}

func ErrorHandler(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {
	// return Internal when Marshal failed
	var msg = "internal error"
	var customStatus *runtime.HTTPStatusError
	if errors.As(err, &customStatus) {
		err = customStatus.Err
	}
	s := status.Convert(err)
	pb := s.Proto()

	w.Header().Del("Trailer")
	w.Header().Del("Transfer-Encoding")

	contentType := marshaler.ContentType(pb)
	w.Header().Set("Content-Type", contentType)

	code := campusapis.UnwarpCode(s.Code())

	if code.GrpcCode == codes.Unauthenticated {
		w.Header().Set("WWW-Authenticate", s.Message())
	}

	st := runtime.HTTPStatusFromCode(code.GrpcCode)
	if customStatus != nil {
		st = customStatus.HTTPStatus
	} else if code.GrpcCode == codes.Unknown && pb.Message != "" {
		st = http.StatusOK
	}
	//当httpStatus为非200且错误码为0时，错误码设为httpStatus*100
	codeStatus := lo.Ternary(
		code.Status == campusapis.Status_OK && st != http.StatusOK,
		campusapis.Status(st*100),
		code.Status,
	)
	switch {
	case pb.Message != "" && codeStatus == campusapis.Status_OK:
		msg = pb.Message
	case pb.Message != "" && codeStatus != campusapis.Status_OK:
		msg = codeStatus.Message() + ": " + pb.Message
	case pb.Message == "" && codeStatus == campusapis.Status_OK:
		msg = "success"
	default:
		msg = codeStatus.Message()
	}
	if mw, ok := w.(*bodyWarper); ok {
		mw.code = lo.Ternary(codeStatus == campusapis.Status_OK, 0, int(codeStatus))
		mw.message = msg
	}

	w.WriteHeader(st)
}

type bodyWarper struct {
	code    int
	message string
	// noWarpFlag 标记是否需要包一层标准返回
	noWarpFlag bool
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyWarper) Write(b []byte) (int, error) {
	return w.body.Write(b)
}

var headerAllowList = make([]string, 0)

func AddToHeaderAllowList(allowList ...string) {
	for _, allowed := range allowList {
		headerAllowList = append(headerAllowList, textproto.CanonicalMIMEHeaderKey(allowed))
	}
}

func HeaderWarp(key string) (string, bool) {
	key = textproto.CanonicalMIMEHeaderKey(key)
	for _, allowed := range headerAllowList {
		if key == allowed {
			return key, true
		}
	}
	return runtime.DefaultHeaderMatcher(key)
}

func GinResponseWarpMiddleware(c *gin.Context) {
	blw := &bodyWarper{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	originalWriter := c.Writer
	c.Writer = blw

	c.Next()

	var responseBytes []byte
	var err error
	if !blw.noWarpFlag {
		responseBytes, err = json.Marshal(&response{
			Code: blw.code,
			Msg:  lo.Ternary(blw.message == "", "success", blw.message),
			//处理如果返回的是空，则替换为null
			Data: json.RawMessage(lo.Ternary(
				len(blw.body.Bytes()) == 0,
				[]byte("null"),
				blw.body.Bytes(),
			)),
		})
		if err != nil {
			responseBytes = []byte(`{"code":50000,"msg":"internal error"}`)
		}
	} else {
		responseBytes = blw.body.Bytes()
	}

	originalWriter.Write(responseBytes)
	c.Abort()
}

type response struct {
	Data json.RawMessage `json:"data"`
	Code int             `json:"error"`
	Msg  string          `json:"msg"`
}
