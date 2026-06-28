package logic

import (
	"context"

	"demo4mysql/internal/svc"
	"demo4mysql/user_mgr_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUserLogic) DeleteUser(in *user_mgr_pb.DeleteUserReq) (*user_mgr_pb.DeleteUserRsp, error) {
	// todo: add your logic here and delete this line
	err := l.svcCtx.TUserModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &user_mgr_pb.DeleteUserRsp{}, nil
}
