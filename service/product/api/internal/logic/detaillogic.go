package logic

import (
	"context"
	"go-zero-mall/service/product/rpc/product"

	"go-zero-mall/service/product/api/internal/svc"
	"go-zero-mall/service/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.DetailRequest) (resp *types.DetailResponse, err error) {
	product, err := l.svcCtx.ProductRpc.Detail(l.ctx, &product.DetailRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &types.DetailResponse{
		Id:     product.Id,
		Name:   product.Name,
		Desc:   product.Desc,
		Stock:  product.Stock,
		Amount: product.Amount,
		Status: product.Status,
	}, nil
}
