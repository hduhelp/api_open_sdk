package grpcgateway

import (
	"net"

	"github.com/soheilhy/cmux"
)

func GRPCGatewayListener(mux cmux.CMux) net.Listener {
	return mux.Match(cmux.HTTP1Fast())
}

func GRPCServerListener(mux cmux.CMux) net.Listener {
	return mux.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
}
