// Code generated by goctl. DO NOT EDIT.
// versions:
//  goctl version: 1.7.3

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	payFieldNames          = builder.RawFieldNames(&Pay{})
	payRows                = strings.Join(payFieldNames, ",")
	payRowsExpectAutoSet   = strings.Join(stringx.Remove(payFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	payRowsWithPlaceHolder = strings.Join(stringx.Remove(payFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cachePayIdPrefix = "cache:pay:id:"
)

type (
	payModel interface {
		Insert(ctx context.Context, data *Pay) (sql.Result, error)
		FindOne(ctx context.Context, id uint64) (*Pay, error)
		Update(ctx context.Context, data *Pay) error
		Delete(ctx context.Context, id uint64) error
	}

	defaultPayModel struct {
		sqlc.CachedConn
		table string
	}

	Pay struct {
		Id         uint64    `db:"id"`
		Uid        uint64    `db:"uid"`    // 用户ID
		Oid        uint64    `db:"oid"`    // 订单ID
		Amount     uint64    `db:"amount"` // 产品金额
		Source     uint64    `db:"source"` // 支付方式
		Status     uint64    `db:"status"` // 支付状态
		CreateTime time.Time `db:"create_time"`
		UpdateTime time.Time `db:"update_time"`
	}
)

func newPayModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultPayModel {
	return &defaultPayModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`pay`",
	}
}

func (m *defaultPayModel) Delete(ctx context.Context, id uint64) error {
	payIdKey := fmt.Sprintf("%s%v", cachePayIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, payIdKey)
	return err
}

func (m *defaultPayModel) FindOne(ctx context.Context, id uint64) (*Pay, error) {
	payIdKey := fmt.Sprintf("%s%v", cachePayIdPrefix, id)
	var resp Pay
	err := m.QueryRowCtx(ctx, &resp, payIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", payRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultPayModel) Insert(ctx context.Context, data *Pay) (sql.Result, error) {
	payIdKey := fmt.Sprintf("%s%v", cachePayIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, payRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Uid, data.Oid, data.Amount, data.Source, data.Status)
	}, payIdKey)
	return ret, err
}

func (m *defaultPayModel) Update(ctx context.Context, data *Pay) error {
	payIdKey := fmt.Sprintf("%s%v", cachePayIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, payRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Uid, data.Oid, data.Amount, data.Source, data.Status, data.Id)
	}, payIdKey)
	return err
}

func (m *defaultPayModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cachePayIdPrefix, primary)
}

func (m *defaultPayModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", payRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultPayModel) tableName() string {
	return m.table
}
