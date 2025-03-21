package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-mall/service/order/api/internal/config"
	"go-zero-mall/service/order/rpc/order"
	"go-zero-mall/service/product/rpc/product"
)

type ServiceContext struct {
	Config     config.Config
	OrderRpc   order.Order
	ProductRpc product.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		OrderRpc:   order.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
		ProductRpc: product.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
	}
}
