package logic

import (
	"context"
	"go-zero-mall/service/product/model"
	"google.golang.org/grpc/status"

	"go-zero-mall/service/product/rpc/internal/svc"
	"go-zero-mall/service/product/rpc/pb"

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
	newProduct := model.Product{
		Name:   in.Name,
		Desc:   in.Desc,
		Stock:  in.Stock,
		Status: in.Status,
		Amount: in.Amount,
	}

	res, err := l.svcCtx.ProductModel.Insert(l.ctx, &newProduct)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	newProductId, err := res.LastInsertId()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	newProduct.Id = uint64(newProductId)
	return &pb.CreateResponse{
		Id: newProduct.Id,
	}, nil
}
