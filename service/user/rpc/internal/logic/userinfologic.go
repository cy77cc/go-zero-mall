package logic

import (
	"context"
	"errors"
	"go-zero-mall/service/user/model"
	"google.golang.org/grpc/status"

	"go-zero-mall/service/user/rpc/internal/svc"
	"go-zero-mall/service/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *pb.UserInfoRequest) (*pb.UserInfoResponse, error) {
	// todo: add your logic here and delete this line
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)

	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Error(100, "用户不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	return &pb.UserInfoResponse{
		Id:     user.Id,
		Name:   user.Name,
		Gender: user.Gender,
		Mobile: user.Mobile,
	}, nil
}
