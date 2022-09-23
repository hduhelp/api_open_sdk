package grpcgateway

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/textproto"
	"strconv"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/hduhelp/api_open_sdk/common"
	"github.com/samber/lo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Register func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error

func DefaultRoutingErrorHandler(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, r *http.Request, httpStatus int) {
	sterr := status.Error(codes.Internal, "Unexpected routing error")
	switch httpStatus {
	case http.StatusBadRequest:
		sterr = status.Error(codes.InvalidArgument, "")
	case http.StatusMethodNotAllowed:
		sterr = status.Error(codes.Unimplemented, "")
	case http.StatusNotFound:
		sterr = status.Error(codes.NotFound, "router not found")
	}
	DefaultErrorHandler(ctx, mux, marshaler, w, r, sterr)
}

func DefaultErrorHandler(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {
	var msg string
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

	code := common.UnwarpCode(s.Code())

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
		code.Status == common.Status_OK && st != http.StatusOK,
		common.Status(st*100),
		code.Status,
	)
	switch {
	case pb.Message != "" && codeStatus == common.Status_OK:
		msg = pb.Message
	case pb.Message != "" && codeStatus != common.Status_OK:
		msg = codeStatus.Message() + ": " + pb.Message
	case pb.Message == "" && codeStatus == common.Status_OK:
		msg = "success"
	default:
		msg = codeStatus.Message()
	}
	if mw, ok := w.(*ResponseWriter); ok {
		mw.code = lo.Ternary(codeStatus == common.Status_OK, 0, int(codeStatus))
		mw.message = msg
	} else {
		msg = strings.ReplaceAll(msg, "\\", "\\\\")
		msg = strings.ReplaceAll(msg, "\"", "\\\"")
		w.Write([]byte("{\"error\": " + lo.Ternary(codeStatus == common.Status_OK, "0", strconv.Itoa(int(codeStatus))) + ", \"msg\": \"" + msg + "\"}"))
	}

	w.WriteHeader(st)
}

var headerAllowList = make([]string, 0)

func init() {
	AddToHeaderAllowList(
		"sign",
		"requestid",
		"userid",
		"staffid",
		"stafftype",
	)
}

func AddToHeaderAllowList(allowList ...string) {
	for _, allowed := range allowList {
		headerAllowList = append(headerAllowList, textproto.CanonicalMIMEHeaderKey(allowed))
	}
}

func DefaultHeaderWarp(key string) (string, bool) {
	key = textproto.CanonicalMIMEHeaderKey(key)
	for _, allowed := range headerAllowList {
		if key == allowed {
			return key, true
		}
	}
	return runtime.DefaultHeaderMatcher(key)
}

type response struct {
	Data json.RawMessage `json:"data"`
	Code int             `json:"error"`
	Msg  string          `json:"msg"`
}

type ResponseWriter struct {
	code    int
	message string
	// noWarpFlag 标记是否需要包一层标准返回
	noWarpFlag bool
	http.ResponseWriter
}

func (w ResponseWriter) Write(body []byte) (int, error) {
	switch {
	case len(body) == 0:
		body = []byte("null")
	case w.noWarpFlag:
		return w.ResponseWriter.Write(body)
	}
	resp := &response{
		Code: w.code,
		Msg:  lo.Ternary(w.message == "", "success", w.message),
		Data: json.RawMessage(body),
	}
	responseBytes, err := json.Marshal(resp)
	switch {
	case err != nil && w.message == "":
		return w.ResponseWriter.Write(body)
	case err != nil && w.message != "":
		resp.Data, _ = json.Marshal(string(body))
		return w.ResponseWriter.Write(lo.Must(json.Marshal(resp)))
	}
	return w.ResponseWriter.Write(responseBytes)
}

func WithResponseWriter(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		handler.ServeHTTP(ResponseWriter{ResponseWriter: writer}, request)
	})
}
