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

	product, err := l.svcCtx.ProductModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Error(100, "产品不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	err = l.svcCtx.ProductModel.Delete(l.ctx, product.Id)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	return &pb.RemoveResponse{}, nil
}
