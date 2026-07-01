package logic

import (
	"context"
	"errors"

	"demo5mysql/internal/svc"
	"demo5mysql/user_mgr_pb"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
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
	// 校验字段长度
	// UserId[1,64]
	if len(in.UserId) < 1 || len(in.UserId) > 64 {
		return nil, errors.New("userId length is not in range [1,64]")
	}

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
