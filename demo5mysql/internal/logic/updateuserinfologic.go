package logic

import (
	"context"
	"errors"

	"demo5mysql/internal/svc"
	"demo5mysql/model/mysql"
	"demo5mysql/user_mgr_pb"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
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
	// 1. 校验用户是否存在
	relation, err := l.svcCtx.TRelationModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, errors.New("user not registered")
		} else {
			return nil, err
		}
	}
	if relation.State != RelationStateRegistered {
		return nil, errors.New("user not registered")
	}

	// 2. 更新用户信息
	l.svcCtx.TUserInfoModel.Update(l.ctx, &mysql.TUserInfo{
		Uid:  relation.Uid,
		Name: in.Name,
		Age:  int64(in.Age),
	})

	// 3. 返回更新后的用户信息
	return &user_mgr_pb.UpdateUserInfoRsp{
		UserId: relation.UserId,
	}, nil
}
