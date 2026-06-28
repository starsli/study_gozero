package logic

import (
	"context"

	"demo4mysql/internal/svc"
	"demo4mysql/user_mgr_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllUsersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllUsersLogic {
	return &GetAllUsersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAllUsersLogic) GetAllUsers(in *user_mgr_pb.GetAllUsersReq) (*user_mgr_pb.GetAllUsersRsp, error) {
	users, err := l.svcCtx.TUserModel.FindAll(l.ctx)
	if err != nil {
		return nil, err
	}
	resp := make([]*user_mgr_pb.GetUserRsp, 0, len(users))
	for _, user := range users {
		resp = append(resp, &user_mgr_pb.GetUserRsp{
			Id:   user.Id,
			Name: user.Name,
			Age:  int32(user.Age),
		})
	}
	return &user_mgr_pb.GetAllUsersRsp{
		Users: resp,
	}, nil
}
