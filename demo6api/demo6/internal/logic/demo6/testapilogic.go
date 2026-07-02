// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package demo6

import (
	"context"

	"demo6/internal/svc"
	"demo6/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestapiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 测试api
func NewTestapiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestapiLogic {
	return &TestapiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TestapiLogic) Testapi(req *types.TestApiReq) (resp *types.TestApiRsp, err error) {
	resp = &types.TestApiRsp{
		OutputParams: "hello:" + req.InputParams,
	}
	err = nil
	return
}
