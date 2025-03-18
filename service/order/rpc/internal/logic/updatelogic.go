package logic

import (
	"context"
	"errors"
	"go-zero-mall/service/order/model"
	"google.golang.org/grpc/status"

	"go-zero-mall/service/order/rpc/internal/svc"
	"go-zero-mall/service/order/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLogic) Update(in *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	// 先查询订单是否存在
	res, err := l.svcCtx.OrderModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Error(100, "订单不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	if in.Uid != 0 {
		res.Uid = in.Uid
	}
	if in.Pid != 0 {
		res.Pid = in.Pid
	}
	if in.Amount != 0 {
		res.Amount = in.Amount
	}
	if in.Status != 0 {
		res.Status = in.Status
	}
	err = l.svcCtx.OrderModel.Update(l.ctx, res)

	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &pb.UpdateResponse{}, nil
}
