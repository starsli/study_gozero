package logic

import (
	"context"

	"demo4mysql/internal/svc"
	"demo4mysql/model/mysql"
	"demo4mysql/user_mgr_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserLogic) UpdateUser(in *user_mgr_pb.UpdateUserReq) (*user_mgr_pb.UpdateUserRsp, error) {
	user := &mysql.TUser{
		Id:   in.Id,
		Name: in.Name,
		Age:  int64(in.Age),
	}
	err := l.svcCtx.TUserModel.Update(l.ctx, user)
	if err != nil {
		return nil, err
	}
	return &user_mgr_pb.UpdateUserRsp{}, nil
}
