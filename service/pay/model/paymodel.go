package model

import (
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var (
	_                 PayModel = (*customPayModel)(nil)
	cachePayOidPrefix          = "cache:pay:oid:"
)

type (
	// PayModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPayModel.
	PayModel interface {
		payModel
		FindOneByOid(oid uint64) (*Pay, error)
	}

	customPayModel struct {
		*defaultPayModel
	}
)

// NewPayModel returns a model for the database table.
func NewPayModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) PayModel {
	return &customPayModel{
		defaultPayModel: newPayModel(conn, c, opts...),
	}
}

func (m *defaultPayModel) FindOneByOid(oid uint64) (*Pay, error) {
	payOidKey := fmt.Sprintf("%s%v", cachePayOidPrefix, oid)
	var resp Pay
	err := m.QueryRow(&resp, payOidKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `oid` = ? limit 1", payRows, m.table)
		return conn.QueryRow(v, query, oid)
	})
	switch {
	case err == nil:
		return &resp, nil
	case errors.Is(err, sqlc.ErrNotFound):
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
