package logic

import (
	"context"

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
	if err := CheckUserId(in.UserId); err != nil {
		return nil, err
	}

	relation, err := CheckUserRegistered(l.ctx, l.svcCtx.TRelationModel, in.UserId)
	if err != nil {
		return nil, err
	}

	account, err := l.svcCtx.TAccountModel.FindOne(l.ctx, relation.Uid)
	if err != nil {
		return nil, err
	}

	return &user_mgr_pb.GetUserBalanceRsp{
		UserId:  account.UserId,
		Balance: account.Balance,
	}, nil
}
