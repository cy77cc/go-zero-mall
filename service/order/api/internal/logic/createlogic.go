package logic

import (
	"context"
	"github.com/dtm-labs/dtmgrpc"
	"go-zero-mall/service/order/rpc/order"
	"go-zero-mall/service/product/rpc/product"
	"google.golang.org/grpc/status"

	"go-zero-mall/service/order/api/internal/svc"
	"go-zero-mall/service/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.CreateRequest) (resp *types.CreateResponse, err error) {
	// 获取OrderRpc BuildTarget
	orderRpcBusiServer, err := l.svcCtx.Config.OrderRpc.BuildTarget()
	if err != nil {
		return nil, status.Error(500, "订单创建异常")
	}

	// 获取 ProductRpc BuildTarget
	productRpcBusiServer, err := l.svcCtx.Config.ProductRpc.BuildTarget()
	if err != nil {
		return nil, status.Error(100, "订单创建异常")
	}

	// dtm服务的etcd注册地址
	var dtmServer = "etcd://etcd:2379/dtmservice"
	// 创建一个gid
	gid := dtmgrpc.MustGenGid(dtmServer)
	// 创建一个saga协议的事务
	saga := dtmgrpc.NewSagaGrpc(dtmServer, gid).
		Add(orderRpcBusiServer+"/order.Order/Create",
			orderRpcBusiServer+"/order.Order/CreateRevert", &order.CreateRequest{
				Uid:    req.Uid,
				Pid:    req.Pid,
				Amount: req.Amount,
				Status: 0,
			}).
		Add(productRpcBusiServer+"/product.Product/DecrStock",
			productRpcBusiServer+"/product.Product/DecrStockRevert", &product.DecrStockRequest{
				Id:  req.Pid,
				Num: 1,
			})

	saga.TimeoutToFail = 1800

	err = saga.Submit()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	res, err := l.svcCtx.OrderRpc.Create(l.ctx, &order.CreateRequest{
		Uid:    req.Uid,
		Pid:    req.Pid,
		Amount: req.Amount,
		Status: req.Status,
	})

	if err != nil {
		return nil, err
	}

	resp.Id = res.Id
	return
}
