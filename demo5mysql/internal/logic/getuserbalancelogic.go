package logic

import (
	"context"
	"errors"

	"demo5mysql/internal/svc"
	"demo5mysql/user_mgr_pb"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type GetUserBalanceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserBalanceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserBalanceLogic {
	return &GetUserBalanceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserBalanceLogic) GetUserBalance(in *user_mgr_pb.GetUserBalanceReq) (*user_mgr_pb.GetUserBalanceRsp, error) {
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

	l.Logger.Infof("starsli relation: %v", relation)
	account, err := l.svcCtx.TAccountModel.FindOne(l.ctx, relation.Uid)
	if err != nil {
		return nil, err
	}

	l.Logger.Infof("starsli account: %v", account)
	return &user_mgr_pb.GetUserBalanceRsp{
		UserId:  account.UserId,
		Balance: account.Balance,
	}, nil
}
