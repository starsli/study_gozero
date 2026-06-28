package svc

import (
	"demo4mysql/internal/config"
	"demo4mysql/model/mysql"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config     config.Config
	TUserModel mysql.TUserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		TUserModel: mysql.NewTUserModel(sqlx.NewMysql(c.DB.DataSource),
			cache.CacheConf(c.CacheConf)),
	}
}
