package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-mall/service/order/rpc/order"
	"go-zero-mall/service/pay/model"
	"go-zero-mall/service/pay/rpc/internal/config"
	"go-zero-mall/service/user/rpc/user"
)

type ServiceContext struct {
	Config   config.Config
	PayModel model.PayModel
	UserRpc  user.User
	OrderRpc order.Order
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:   c,
		PayModel: model.NewPayModel(conn, c.CacheRedis),
		UserRpc:  user.NewUser(zrpc.MustNewClient(c.UserRpc)),
		OrderRpc: order.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
	}
}
