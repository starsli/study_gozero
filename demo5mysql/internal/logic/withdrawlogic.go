package logic

import (
	"context"
	"errors"

	"demo5mysql/internal/svc"
	"demo5mysql/model/mysql"
	"demo5mysql/user_mgr_pb"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
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
	if err := CheckUserId(in.UserId); err != nil {
		return nil, err
	}
	if err := CheckAmount(in.Amount); err != nil {
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
	if userInfo.Password != in.Password {
		return nil, errors.New("password is incorrect")
	}

	flowId := uuid.New().String()
	var retBalance int64 = 0
	err = l.svcCtx.SqlConn.TransactCtx(l.ctx, func(ctx context.Context, tx sqlx.Session) error {
		txAccount := l.svcCtx.TAccountModel.WithSession(tx)
		txFlow := l.svcCtx.TAccountFlowModel.WithSession(tx)

		account, err := txAccount.FindOneForUpdate(ctx, relation.Uid)
		if err != nil {
			return err
		}

		retBalance = account.Balance - in.Amount
		if retBalance < 0 {
			return errors.New("balance is not enough")
		}

		account.Balance = retBalance
		err = txAccount.Update(ctx, account)
		if err != nil {
			return err
		}

		_, err = txFlow.Insert(ctx, &mysql.TAccountFlow{
			Uid:            relation.Uid,
			UserId:         in.UserId,
			FlowId:         flowId,
			FlowType:       FlowTypeOut,
			BizType:        BizTypeWithdraw,
			Amount:         in.Amount,
			CounterpartyId: in.BankType,
		})

		// TODO 记录提现单

		return err
	})
	if err != nil {
		return nil, err
	}
	return &user_mgr_pb.WithdrawRsp{
		UserId:  in.UserId,
		Balance: retBalance,
	}, nil
}
