package logic

import (
	"context"
	"database/sql"
	"errors"
	"github.com/dtm-labs/dtmgrpc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-mall/service/order/model"
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
	// 获取rawDB
	db, err := sqlx.NewMysql(l.svcCtx.Config.Mysql.DataSource).RawDB()

	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	// 获取子事务屏障对象
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		return nil, status.Error(500, err.Error())

	}

	newOrder := model.Order{
		Uid:    in.Uid,
		Pid:    in.Pid,
		Amount: in.Amount,
		Status: 0,
	}

	// 开启子事务屏障
	if err := barrier.CallWithDB(db, func(tx *sql.Tx) error {
		//	查询用户是否存在
		_, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
			Id: in.Uid,
		})
		if err != nil {
			return errors.New("用户不存在")
		}

		//	创建订单
		res, err := l.svcCtx.OrderModel.TxInsert(tx, &newOrder)

		id, err := res.LastInsertId()
		if err != nil {
			return errors.New("订单创建失败")
		}

		newOrder.Id = uint64(id)
		return nil
	}); err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &pb.CreateResponse{
		Id: newOrder.Id,
	}, nil
}
