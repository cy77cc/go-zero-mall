package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-mall/service/product/api/internal/config"
	"go-zero-mall/service/product/rpc/product"
)

type ServiceContext struct {
	Config     config.Config
	ProductRpc product.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		ProductRpc: product.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
	}
}
