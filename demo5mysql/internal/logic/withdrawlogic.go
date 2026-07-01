package logic

import (
	"context"
	"errors"

	"demo5mysql/internal/svc"
	"demo5mysql/model/mysql"
	"demo5mysql/user_mgr_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type WithdrawLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewWithdrawLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WithdrawLogic {
	return &WithdrawLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *WithdrawLogic) Withdraw(in *user_mgr_pb.WithdrawReq) (*user_mgr_pb.WithdrawRsp, error) {
	relation, err := l.svcCtx.TRelationModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}
	if relation.State != 2 {
		return nil, errors.New("user not registered")
	}

	var retBalance int64 = 0
	err = l.svcCtx.TAccountModel.TransactCtx(l.ctx, func(ctx context.Context, tx mysql.TAccountModel) error {
		account, err := tx.FindOneForUpdate(ctx, relation.Uid)
		if err != nil {
			return err
		}

		retBalance = account.Balance - in.Amount
		account.Balance = retBalance
		return tx.Update(ctx, account)
	})
	if err != nil {
		return nil, err
	}
	return &user_mgr_pb.WithdrawRsp{
		UserId:  in.UserId,
		Balance: retBalance,
	}, nil
}
