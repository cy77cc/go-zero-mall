package model

import (
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ProductModel = (*customProductModel)(nil)

type (
	// ProductModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProductModel.
	ProductModel interface {
		productModel
		TxAdjustStock(tx *sql.Tx, id uint64, delta int) (sql.Result, error)
	}

	customProductModel struct {
		*defaultProductModel
	}
)

// NewProductModel returns a model for the database table.
func NewProductModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ProductModel {
	return &customProductModel{
		defaultProductModel: newProductModel(conn, c, opts...),
	}
}

func (m *defaultProductModel) TxAdjustStock(tx *sql.Tx, id uint64, delta int) (sql.Result, error) {
	productIdKey := fmt.Sprintf("%s%v", cacheProductIdPrefix, id)
	return m.Exec(
		func(conn sqlx.SqlConn) (result sql.Result, err error) {
			query := fmt.Sprintf("update %s set stock=stock+? where stock >= -? and id=?", m.table)
			return tx.Exec(query, delta, delta, id)
		},
		productIdKey,
	)
}
