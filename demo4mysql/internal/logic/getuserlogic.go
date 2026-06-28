package logic

import (
	"context"

	"demo4mysql/internal/svc"
	"demo4mysql/user_mgr_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user_mgr_pb.GetUserReq) (*user_mgr_pb.GetUserRsp, error) {
	user, err := l.svcCtx.TUserModel.FindOneByID(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &user_mgr_pb.GetUserRsp{
		Id:   user.Id,
		Name: user.Name,
		Age:  int32(user.Age),
	}, nil
}
