package logic

import (
	"context"
	"errors"
	"go-zero-mall/service/order/rpc/order"
	"go-zero-mall/service/pay/model"
	"go-zero-mall/service/user/rpc/user"
	"google.golang.org/grpc/status"

	"go-zero-mall/service/pay/rpc/internal/svc"
	"go-zero-mall/service/pay/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CallbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CallbackLogic {
	return &CallbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CallbackLogic) Callback(in *pb.CallbackRequest) (*pb.CallbackResponse, error) {
	// 查询用户是否存在
	_, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
		Id: in.Id,
	})

	if err != nil {
		return nil, err
	}

	// 查询订单是否存在
	_, err = l.svcCtx.OrderRpc.Detail(l.ctx, &order.DetailRequest{
		Id: in.Oid,
	})

	if err != nil {
		return nil, err
	}

	//查询支付是否存在

	res, err := l.svcCtx.PayModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Error(100, "支付不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	// 支付金额与订单金额不符
	if in.Amount != res.Amount {
		return nil, status.Error(100, "支付金额与订单金额不符")
	}

	res.Source = in.Source
	res.Status = in.Status

	err = l.svcCtx.PayModel.Update(l.ctx, res)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	// 修改订单支付状态
	_, err = l.svcCtx.OrderRpc.Paid(l.ctx, &order.PaidRequest{
		Id: in.Oid,
	})

	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &pb.CallbackResponse{}, nil
}
