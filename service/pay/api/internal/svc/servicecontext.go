package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-mall/service/pay/api/internal/config"
	"go-zero-mall/service/pay/rpc/pay"
)

type ServiceContext struct {
	Config config.Config
	PayRpc pay.Pay
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		PayRpc: pay.NewPay(zrpc.MustNewClient(c.PayRpc)),
	}
}
