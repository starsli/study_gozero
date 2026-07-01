package logic

import (
	"context"

	"demo5mysql/internal/svc"
	"demo5mysql/model/mysql"
	"demo5mysql/user_mgr_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserInfoLogic {
	return &UpdateUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserInfoLogic) UpdateUserInfo(in *user_mgr_pb.UpdateUserInfoReq) (*user_mgr_pb.UpdateUserInfoRsp, error) {
	if err := CheckUserId(in.UserId); err != nil {
		return nil, err
	}

	relation, err := CheckUserRegistered(l.ctx, l.svcCtx.TRelationModel, in.UserId)
	if err != nil {
		return nil, err
	}

	// 先查询用户信息，在事务中更新
	err = l.svcCtx.TUserInfoModel.TransactCtx(l.ctx, func(ctx context.Context, tx mysql.TUserInfoModel) error {
		userInfo, err := tx.FindOne(ctx, relation.Uid)
		if err != nil {
			return err
		}

		userInfo.Name = in.Name
		userInfo.Age = int64(in.Age)
		userInfo.Gender = int64(in.Gender)
		userInfo.Address = in.Address
		userInfo.Phone = in.Phone
		userInfo.Email = in.Email
		userInfo.IdType = int64(in.IdType)
		userInfo.IdCard = in.IdCard

		return tx.Update(ctx, userInfo)
	})
	if err != nil {
		return nil, err
	}

	return &user_mgr_pb.UpdateUserInfoRsp{
		UserId: relation.UserId,
	}, nil
}
