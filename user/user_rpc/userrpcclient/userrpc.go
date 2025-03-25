// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.1
// Source: user_rpc.proto

package userrpcclient

import (
	"context"

	"Easy-GoZero/user/user_rpc/user_rpc"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Request  = user_rpc.Request
	Response = user_rpc.Response

	UserRpc interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	}

	defaultUserRpc struct {
		cli zrpc.Client
	}
)

func NewUserRpc(cli zrpc.Client) UserRpc {
	return &defaultUserRpc{
		cli: cli,
	}
}

func (m *defaultUserRpc) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := user_rpc.NewUserRpcClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}
