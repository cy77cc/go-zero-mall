package logic

import (
	"context"
	"errors"
	"go-zero-mall/service/product/model"
	"google.golang.org/grpc/status"

	"go-zero-mall/service/product/rpc/internal/svc"
	"go-zero-mall/service/product/rpc/pb"

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
	// 查询产品是否存在
	product, err := l.svcCtx.ProductModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Error(100, "产品不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	if in.Name != "" {
		product.Name = in.Name
	}
	if in.Desc != "" {
		product.Desc = in.Desc
	}
	if in.Stock != 0 {
		product.Stock = in.Stock
	}
	if in.Amount != 0 {
		product.Amount = in.Amount
	}
	if in.Status != 0 {
		product.Status = in.Status
	}

	err = l.svcCtx.ProductModel.Update(l.ctx, product)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &pb.UpdateResponse{}, nil
}
