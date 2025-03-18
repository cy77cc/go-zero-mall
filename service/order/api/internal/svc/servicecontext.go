package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-mall/service/order/api/internal/config"
	"go-zero-mall/service/order/rpc/order"
)

type ServiceContext struct {
	Config   config.Config
	OrderRpc order.Order
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		OrderRpc: order.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
	}
}
