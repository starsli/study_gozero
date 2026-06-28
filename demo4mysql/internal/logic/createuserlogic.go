package logic

import (
	"context"

	"demo4mysql/internal/svc"
	"demo4mysql/model/mysql"
	"demo4mysql/user_mgr_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserLogic) CreateUser(in *user_mgr_pb.CreateUserReq) (*user_mgr_pb.CreateUserRsp, error) {
	user := &mysql.TUser{
		Name: in.Name,
		Age:  int64(in.Age),
	}
	_, err := l.svcCtx.TUserModel.Insert(l.ctx, user)
	if err != nil {
		return nil, err
	}

	return &user_mgr_pb.CreateUserRsp{}, nil
}
