package mysql

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TAccountFlowModel = (*customTAccountFlowModel)(nil)

type (
	// TAccountFlowModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTAccountFlowModel.
	TAccountFlowModel interface {
		tAccountFlowModel
		WithSession(session sqlx.Session) TAccountFlowModel
		// FindByUidWithOrderByTime 根据uid查询用户流水, 根据create_time排序, 从新到旧查询
		FindByUidWithOrderByTime(ctx context.Context, uid int64, offset, limit int32) ([]*TAccountFlow, error)
	}

	customTAccountFlowModel struct {
		*defaultTAccountFlowModel
	}
)

// NewTAccountFlowModel returns a model for the database table.
func NewTAccountFlowModel(conn sqlx.SqlConn) TAccountFlowModel {
	return &customTAccountFlowModel{
		defaultTAccountFlowModel: newTAccountFlowModel(conn),
	}
}

func (m *customTAccountFlowModel) WithSession(session sqlx.Session) TAccountFlowModel {
	return NewTAccountFlowModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customTAccountFlowModel) FindByUid(ctx context.Context, uid int64) ([]*TAccountFlow, error) {
	query := "select " + tAccountFlowRows + " from " + m.table + " where `uid` = ? order by `create_time` desc"
	var resp []*TAccountFlow
	err := m.conn.QueryRowsCtx(ctx, &resp, query, uid)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return nil, nil
	default:
		return nil, err
	}
}

func (m *customTAccountFlowModel) FindByUidWithOrderByTime(ctx context.Context, uid int64, offset, limit int32) ([]*TAccountFlow, error) {
	query := "select " + tAccountFlowRows + " from " + m.table + " where `uid` = ? order by `create_time` desc limit ?, ?"
	var resp []*TAccountFlow
	err := m.conn.QueryRowsCtx(ctx, &resp, query, uid, offset, limit)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return nil, nil
	default:
		return nil, err
	}
}
