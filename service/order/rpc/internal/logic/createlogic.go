package logic

import (
	"context"
	"go-zero-mall/service/order/model"
	"go-zero-mall/service/product/rpc/product"
	"go-zero-mall/service/user/rpc/user"
	"google.golang.org/grpc/status"

	"go-zero-mall/service/order/rpc/internal/svc"
	"go-zero-mall/service/order/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *pb.CreateRequest) (*pb.CreateResponse, error) {
	_, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
		Id: in.Uid,
	})

	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	productRes, err := l.svcCtx.ProductRpc.Detail(l.ctx, &product.DetailRequest{
		Id: in.Pid,
	})

	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	// 检查产品库存是否充足
	if productRes.Stock <= 0 {
		return nil, status.Error(500, "产品库存不足")
	}

	newOrder := model.Order{
		Uid:    in.Uid,
		Pid:    in.Pid,
		Amount: in.Amount,
		Status: 0,
	}

	res, err := l.svcCtx.OrderModel.Insert(l.ctx, &newOrder)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	newOrder.Id = uint64(id)

	// 更新产品库存
	_, err = l.svcCtx.ProductRpc.Update(l.ctx, &product.UpdateRequest{
		Id:     productRes.Id,
		Name:   productRes.Name,
		Desc:   productRes.Desc,
		Stock:  productRes.Stock - 1,
		Amount: productRes.Amount,
		Status: productRes.Status,
	})

	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &pb.CreateResponse{
		Id: newOrder.Id,
	}, nil
}
