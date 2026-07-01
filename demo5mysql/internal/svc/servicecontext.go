package svc

import (
	"demo5mysql/internal/config"
	"demo5mysql/model/mysql"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config            config.Config
	SqlConn           sqlx.SqlConn
	TRelationModel    mysql.TRelationModel
	TUserInfoModel    mysql.TUserInfoModel
	TAccountModel     mysql.TAccountModel
	TUidSegmentModel  mysql.TUidSegmentModel
	TAccountFlowModel mysql.TAccountFlowModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:            c,
		SqlConn:           sqlConn,
		TRelationModel:    mysql.NewTRelationModel(sqlConn),
		TUserInfoModel:    mysql.NewTUserInfoModel(sqlConn),
		TAccountModel:     mysql.NewTAccountModel(sqlConn),
		TUidSegmentModel:  mysql.NewTUidSegmentModel(sqlConn),
		TAccountFlowModel: mysql.NewTAccountFlowModel(sqlConn),
	}
}
