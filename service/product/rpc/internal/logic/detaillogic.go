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

type DetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DetailLogic) Detail(in *pb.DetailRequest) (*pb.DetailResponse, error) {
	product, err := l.svcCtx.ProductModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Error(100, "产品不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	return &pb.DetailResponse{
		Id:     product.Id,
		Name:   product.Name,
		Desc:   product.Desc,
		Stock:  product.Stock,
		Amount: product.Amount,
		Status: product.Status,
	}, nil
}
