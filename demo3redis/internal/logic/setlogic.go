package logic

import (
	"context"

	"demo3redis/internal/svc"
	demo_pb3 "demo3redis/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetLogic {
	return &SetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetLogic) Set(in *demo_pb3.SetReq) (*demo_pb3.SetRsp, error) {
	key := in.Key
	value := in.Value
	// TODO：处理设置超时的情况
	err := l.svcCtx.Redis.SetCtx(l.ctx, key, value)
	if err != nil {
		return nil, err
	}
	l.Logger.Debugf("starsli Set() key = %s, value = %s", key, value)
	return &demo_pb3.SetRsp{}, nil
}
