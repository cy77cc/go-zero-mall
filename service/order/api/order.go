package main

import (
	"flag"
	"fmt"

	"go-zero-mall/service/order/api/internal/config"
	"go-zero-mall/service/order/api/internal/handler"
	"go-zero-mall/service/order/api/internal/svc"

	_ "github.com/dtm-labs/driver-gozero"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/order.yaml", "the config file")

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
