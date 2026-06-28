package svc

import (
	"demo2/app/rpc1/internal/config"
	"demo2/app/rpc2/rpc2"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	Rpc2   rpc2.Rpc2
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Rpc2:   rpc2.NewRpc2(zrpc.MustNewClient(c.Rpc2Conf)),
	}
}
