package logic

import (
	"context"
	"go-zero-mall/service/product/rpc/product"

	"go-zero-mall/service/product/api/internal/svc"
	"go-zero-mall/service/product/api/internal/types"

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

func (l *CreateLogic) Create(req *types.CreateRequst) (resp *types.CreateResponse, err error) {
	product, err := l.svcCtx.ProductRpc.Create(l.ctx, &product.CreateRequest{
		Name:   req.Name,
		Desc:   req.Desc,
		Stock:  req.Stock,
		Amount: req.Amount,
		Status: req.Status,
	})

	if err != nil {
		return nil, err
	}

	return &types.CreateResponse{
		Id: product.Id,
	}, nil
}
