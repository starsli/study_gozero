package logic

import (
	"context"

	"demo2/app/rpc1/internal/svc"
	"demo2/app/rpc1/rpc1_pb"
	"demo2/app/rpc2/rpc2_pb"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (l *TestCLogic) TestC(in *rpc1_pb.TestCReq) (*rpc1_pb.TestCRsp, error) {
	resp, err := l.svcCtx.Rpc2.TestC(l.ctx, &rpc2_pb.TestCReq{InputParams: in.InputParams})
	if err != nil {
		// 如果是超时错误，只打印日志，不向上报错
		if status.Code(err) == codes.DeadlineExceeded {
			l.Logger.Errorf("rpc2 call timeout: %v", err)
			return &rpc1_pb.TestCRsp{OutputParams: in.InputParams + "starsli1"}, nil
		}
		return nil, err
	}
	return &rpc1_pb.TestCRsp{OutputParams: resp.OutputParams + "starsli1"}, nil
}
