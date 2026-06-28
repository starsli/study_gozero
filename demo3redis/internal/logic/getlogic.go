package logic

import (
	"context"

	"demo3redis/internal/svc"
	demo_pb3 "demo3redis/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogic {
	return &GetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLogic) Get(in *demo_pb3.GetReq) (*demo_pb3.GetRsp, error) {
	key := in.Key
	// TODO：处理获取超时的情况
	value, err := l.svcCtx.Redis.GetCtx(l.ctx, key)
	if err != nil {
		return nil, err
	}
	l.Logger.Debugf("starsli Get() value = %s", value)
	return &demo_pb3.GetRsp{Value: value}, nil
}
