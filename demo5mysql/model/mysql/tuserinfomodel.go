package mysql

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TUserInfoModel = (*customTUserInfoModel)(nil)

type (
	// TUserInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTUserInfoModel.
	TUserInfoModel interface {
		tUserInfoModel
		withSession(session sqlx.Session) TUserInfoModel
	}

	customTUserInfoModel struct {
		*defaultTUserInfoModel
	}
)

// NewTUserInfoModel returns a model for the database table.
func NewTUserInfoModel(conn sqlx.SqlConn) TUserInfoModel {
	return &customTUserInfoModel{
		defaultTUserInfoModel: newTUserInfoModel(conn),
	}
}

func (m *customTUserInfoModel) withSession(session sqlx.Session) TUserInfoModel {
	return NewTUserInfoModel(sqlx.NewSqlConnFromSession(session))
}
