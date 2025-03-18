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

type RemoveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveLogic {
	return &RemoveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveLogic) Remove(in *pb.RemoveRequest) (*pb.RemoveResponse, error) {
	res, err := l.svcCtx.OrderModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Error(100, "订单不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	err = l.svcCtx.OrderModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &pb.RemoveResponse{}, nil
}
