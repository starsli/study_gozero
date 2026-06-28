package logic

import (
	"context"

	"demo1/demo1_pb"
	"demo1/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type Demo1testLogic struct {
	ctx         context.Context
	svcCtx      *svc.ServiceContext
	logx.Logger // 匿名字段，字段名=类型名 Logger
}

func NewDemo1testLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Demo1testLogic {
	return &Demo1testLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *Demo1testLogic) Demo1Test(in *demo1_pb.Demo1Req) (*demo1_pb.Demo2Rsp, error) {
	// todo: add your logic here and delete this line
	OutputParams := "welcome:" + in.InputParams
	l.Logger.Debugf("starsli input_params: %s, output_params: %s", in.InputParams, OutputParams)
	l.Logger.Debugf("starsli1")
	l.Logger.Debugf("starsli2")
	l.Logger.Debugf("starsli3")
	l.Logger.Debugf("starsli4")
	return &demo1_pb.Demo2Rsp{OutputParams: OutputParams}, nil
}
