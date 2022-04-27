package grpcclient

import (
	"context"
	"testing"

	"google.golang.org/grpc"
)

func TestConn(t *testing.T) {
	type args struct {
		ctx       context.Context
		endpoints []string
	}
	tests := []struct {
		name string
		args args
		want grpc.ClientConnInterface
	}{
		{
			name: "with none endpoint",
			args: args{
				ctx:       context.Background(),
				endpoints: []string{},
			},
		},
		{
			name: "with custom endpoint",
			args: args{
				ctx: context.Background(),
				endpoints: []string{
					"gapi.hduhelp.com:443",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Conn(tt.args.ctx, tt.args.endpoints...); got == nil {
				t.Errorf("Conn() = %v", got)
			}
		})
	}
}
