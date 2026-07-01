package mysql

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TAccountModel = (*customTAccountModel)(nil)

type (
	// TAccountModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTAccountModel.
	TAccountModel interface {
		tAccountModel
		WithSession(session sqlx.Session) TAccountModel
		TransactCtx(ctx context.Context, fn func(ctx context.Context, tx TAccountModel) error) error
		FindOneForUpdate(ctx context.Context, uid int64) (*TAccount, error)
	}

	customTAccountModel struct {
		*defaultTAccountModel
	}
)

// NewTAccountModel returns a model for the database table.
func NewTAccountModel(conn sqlx.SqlConn) TAccountModel {
	return &customTAccountModel{
		defaultTAccountModel: newTAccountModel(conn),
	}
}

func (m *customTAccountModel) WithSession(session sqlx.Session) TAccountModel {
	return NewTAccountModel(sqlx.NewSqlConnFromSession(session))
}

// TransactCtx 事务封装
func (m *customTAccountModel) TransactCtx(ctx context.Context, fn func(ctx context.Context, tx TAccountModel) error) error {
	return m.conn.TransactCtx(ctx, func(ctx context.Context, tx sqlx.Session) error {
		return fn(ctx, m.WithSession(tx))
	})
}

// FindOneForUpdate 查询并锁定账户
func (m *customTAccountModel) FindOneForUpdate(ctx context.Context, uid int64) (*TAccount, error) {
	query := "select " + tAccountRows + " from " + m.table + " where `uid` = ? limit 1 for update"
	var resp TAccount
	err := m.conn.QueryRowCtx(ctx, &resp, query, uid)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
