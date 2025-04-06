package main

import (
	"flag"
	"fmt"

	"Easy-GoZero/go-zero_native_operation_mysql_09/user/api/internal/config"
	"Easy-GoZero/go-zero_native_operation_mysql_09/user/api/internal/handler"
	"Easy-GoZero/go-zero_native_operation_mysql_09/user/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/users.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
