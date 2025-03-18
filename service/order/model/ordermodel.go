package model

import (
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OrderModel = (*customOrderModel)(nil)

type (
	// OrderModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOrderModel.
	OrderModel interface {
		orderModel
		FindAllByUid(uid uint64) ([]*Order, error)
	}

	customOrderModel struct {
		*defaultOrderModel
	}
)

// NewOrderModel returns a model for the database table.
func NewOrderModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) OrderModel {
	return &customOrderModel{
		defaultOrderModel: newOrderModel(conn, c, opts...),
	}
}

func (m *defaultOrderModel) FindAllByUid(uid uint64) ([]*Order, error) {
	var list []*Order

	query := fmt.Sprintf("select %s from %s where `uid` = ?", orderRows, m.table)
	err := m.QueryRowNoCache(&list, query, uid)

	switch {
	case err == nil:
		return list, nil
	case errors.Is(err, sqlx.ErrNotFound):
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
