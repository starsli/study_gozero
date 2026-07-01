package logic

import (
	"context"
	"errors"

	"demo5mysql/internal/svc"
	"demo5mysql/user_mgr_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *user_mgr_pb.GetUserInfoReq) (*user_mgr_pb.GetUserInfoRsp, error) {
	relation, err := l.svcCtx.TRelationModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}
	if relation.State != 2 {
		return nil, errors.New("user not registered")
	}

	userInfo, err := l.svcCtx.TUserInfoModel.FindOne(l.ctx, relation.Uid)
	if err != nil {
		return nil, err
	}

	return &user_mgr_pb.GetUserInfoRsp{
		UserId: userInfo.UserId,
		Name:   userInfo.Name,
		Age:    int32(userInfo.Age),
	}, nil
}
