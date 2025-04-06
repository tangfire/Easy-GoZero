package main

import (
	"flag"
	"fmt"

	"Easy-GoZero/Service_Grouping_12/user/rpc/internal/config"
	usercreateServer "Easy-GoZero/Service_Grouping_12/user/rpc/internal/server/usercreate"
	userinfoServer "Easy-GoZero/Service_Grouping_12/user/rpc/internal/server/userinfo"
	"Easy-GoZero/Service_Grouping_12/user/rpc/internal/svc"
	"Easy-GoZero/Service_Grouping_12/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserCreateServer(grpcServer, usercreateServer.NewUserCreateServer(ctx))
		user.RegisterUserInfoServer(grpcServer, userinfoServer.NewUserInfoServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
