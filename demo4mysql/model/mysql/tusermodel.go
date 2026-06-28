package mysql

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TUserModel = (*customTUserModel)(nil)

type (
	// TUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTUserModel.
	TUserModel interface {
		tUserModel
		// 添加自定义方法
		FindOneByID(ctx context.Context, id int64) (*TUser, error)
		FindAll(ctx context.Context) ([]*TUser, error)
		FindByAgeRange(ctx context.Context, minAge, maxAge int64) ([]*TUser, error)
	}

	customTUserModel struct {
		*defaultTUserModel
	}
)

// NewTUserModel returns a model for the database table.
func NewTUserModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TUserModel {
	return &customTUserModel{
		defaultTUserModel: newTUserModel(conn, c, opts...),
	}
}

func (m *customTUserModel) FindOneByID(ctx context.Context, id int64) (*TUser, error) {
	var resp TUser
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", tUserRows, m.table)
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customTUserModel) FindAll(ctx context.Context) ([]*TUser, error) {
	var resp []*TUser
	query := fmt.Sprintf("select %s from %s", tUserRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customTUserModel) FindByAgeRange(ctx context.Context, minAge, maxAge int64) ([]*TUser, error) {
	var resp []*TUser
	query := fmt.Sprintf("select %s from %s where `age` >= ? and `age` <= ?", tUserRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, minAge, maxAge)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
