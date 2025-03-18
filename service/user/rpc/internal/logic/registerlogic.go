package logic

import (
	"context"
	"errors"
	"go-zero-mall/common/cryptx"
	"go-zero-mall/service/user/model"
	"google.golang.org/grpc/status"

	"go-zero-mall/service/user/rpc/internal/svc"
	"go-zero-mall/service/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, status.Error(100, "该用户已存在")
	}

	//var newUser model.User

	if errors.Is(err, model.ErrNotFound) {
		newUser := model.User{
			Name:     in.Name,
			Gender:   in.Gender,
			Mobile:   in.Mobile,
			Password: cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password),
		}
		result, err := l.svcCtx.UserModel.Insert(l.ctx, &newUser)

		if err != nil {
			return nil, status.Error(500, err.Error())
		}

		newUserId, err := result.LastInsertId()
		if err != nil {
			return nil, status.Error(500, err.Error())
		}

		newUser.Id = uint64(newUserId)
		return &pb.RegisterResponse{
			Id:     newUser.Id,
			Name:   newUser.Name,
			Gender: int64(newUser.Gender),
			Mobile: newUser.Mobile,
		}, nil
	}

	return nil, status.Error(500, err.Error())

}
