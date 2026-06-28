package logic

import (
	"context"
	"time"

	"demo2/app/rpc2/internal/svc"
	"demo2/app/rpc2/rpc2_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestCLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTestCLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestCLogic {
	return &TestCLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TestCLogic) TestC(in *rpc2_pb.TestCReq) (*rpc2_pb.TestCRsp, error) {
	// todo: add your logic here and delete this line
	l.Logger.Debugf("starsli1")
	time.Sleep(1 * time.Second)
	l.Logger.Debugf("starsli2")
	return &rpc2_pb.TestCRsp{OutputParams: in.InputParams + "starsli2"}, nil
}
