package svc

import (
	"demo5mysql/internal/config"
	"demo5mysql/model/mysql"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config           config.Config
	TRelationModel   mysql.TRelationModel
	TUserInfoModel   mysql.TUserInfoModel
	TAccountModel    mysql.TAccountModel
	TUidSegmentModel mysql.TUidSegmentModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:           c,
		TRelationModel:   mysql.NewTRelationModel(sqlx.NewMysql(c.DB.DataSource)),
		TUserInfoModel:   mysql.NewTUserInfoModel(sqlx.NewMysql(c.DB.DataSource)),
		TAccountModel:    mysql.NewTAccountModel(sqlx.NewMysql(c.DB.DataSource)),
		TUidSegmentModel: mysql.NewTUidSegmentModel(sqlx.NewMysql(c.DB.DataSource)),
	}
}
