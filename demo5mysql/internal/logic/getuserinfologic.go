package logic

import (
	"context"

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
	if err := CheckUserId(in.UserId); err != nil {
		return nil, err
	}

	relation, err := CheckUserRegistered(l.ctx, l.svcCtx.TRelationModel, in.UserId)
	if err != nil {
		return nil, err
	}

	userInfo, err := l.svcCtx.TUserInfoModel.FindOne(l.ctx, relation.Uid)
	if err != nil {
		return nil, err
	}

	return &user_mgr_pb.GetUserInfoRsp{
		UserId:  userInfo.UserId,
		Name:    userInfo.Name,
		Gender:  int32(userInfo.Gender),
		Age:     int32(userInfo.Age),
		Address: userInfo.Address,
		Phone:   userInfo.Phone,
		Email:   userInfo.Email,
		IdType:  int32(userInfo.IdType),
		IdCard:  userInfo.IdCard,
	}, nil
}
