package svc

import (
	"Easy-GoZero/RPC_service_combined_with_API_13/user/api/internal/config"
	"Easy-GoZero/RPC_service_combined_with_API_13/user/rpc/userclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
