package logic

import (
	"context"
	"errors"
	"go-zero-mall/service/order/model"
	"go-zero-mall/service/order/rpc/internal/svc"
	"go-zero-mall/service/order/rpc/order"
	"go-zero-mall/service/order/rpc/pb"
	"go-zero-mall/service/user/rpc/user"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListLogic) List(in *pb.ListRequest) (*pb.ListResponse, error) {
	_, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
		Id: in.Uid,
	})

	if err != nil {
		return nil, err
	}

	//查询订单是否存在
	list, err := l.svcCtx.OrderModel.FindAllByUid(in.Uid)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Error(100, "订单不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	orderList := make([]*order.DetailResponse, 0)

	for _, item := range list {
		orderList = append(orderList, &order.DetailResponse{
			Id:     item.Id,
			Uid:    item.Uid,
			Pid:    item.Pid,
			Amount: item.Amount,
			Status: item.Status,
		})
	}

	return &pb.ListResponse{
		Data: orderList,
	}, nil
}
