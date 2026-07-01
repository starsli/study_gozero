package logic

import (
	"context"
	"errors"

	"demo5mysql/internal/svc"
	"demo5mysql/user_mgr_pb"

	"github.com/zeromicro/go-zero/core/logx"
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
	relation, err := l.svcCtx.TRelationModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}
	if relation.State != 2 {
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
