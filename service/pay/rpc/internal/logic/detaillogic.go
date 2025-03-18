package logic

import (
	"context"
	"errors"
	"go-zero-mall/service/pay/model"
	"google.golang.org/grpc/status"

	"go-zero-mall/service/pay/rpc/internal/svc"
	"go-zero-mall/service/pay/rpc/pb"

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
	res, err := l.svcCtx.PayModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Error(100, "支付不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	return &pb.DetailResponse{
		Id:     res.Id,
		Uid:    res.Uid,
		Oid:    res.Oid,
		Amount: res.Amount,
		Status: res.Status,
		Source: res.Source,
	}, nil
}
