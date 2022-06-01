package campusapis

import (
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (x Status) Warp(status codes.Code) codes.Code {
	return WarpCode(status, uint32(x))
}

func (x Status) With(c codes.Code, msg string) error {
	return status.Error(WarpCode(c, uint32(x)), msg)
}

func (x Status) Message() string {
	return strings.Replace(strings.ToLower(x.String()), "_", " ", -1)
}

func WarpCode(code codes.Code, errorCode uint32) codes.Code {
	return code*100000 + codes.Code(errorCode)
}

type Codes struct {
	Status   Status
	GrpcCode codes.Code
}

func UnwarpCode(code codes.Code) Codes {
	if code < 10000 {
		return Codes{
			Status:   Status_OK,
			GrpcCode: code,
		}
	}
	return Codes{
		Status:   Status(int32(code) % 100000),
		GrpcCode: code / 100000,
	}
}

func Unwarp(err error) (codes Codes, msg string) {
	s := status.Convert(err)
	return UnwarpCode(s.Code()), s.Proto().Message
}
