package logic

import (
	"context"
	"errors"
	"math"

	"demo5mysql/internal/svc"
	"demo5mysql/model/mysql"
	"demo5mysql/user_mgr_pb"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type DepositLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDepositLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DepositLogic {
	return &DepositLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DepositLogic) Deposit(in *user_mgr_pb.DepositReq) (*user_mgr_pb.DepositRsp, error) {
	// 校验字段长度
	// UserId[1,64]
	// Amount[1,int64.Max]
	if len(in.UserId) < 1 || len(in.UserId) > 64 {
		return nil, errors.New("userId length is not in range [1,64]")
	}

	if in.Amount < 1 || in.Amount > math.MaxInt64 {
		return nil, errors.New("amount is not in range [1,int64.Max]")
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

	flowId := uuid.New().String()
	var retBalance int64 = 0
	err = l.svcCtx.SqlConn.TransactCtx(l.ctx, func(ctx context.Context, tx sqlx.Session) error {
		txAccount := l.svcCtx.TAccountModel.WithSession(tx)
		txFlow := l.svcCtx.TAccountFlowModel.WithSession(tx)

		account, err := txAccount.FindOneForUpdate(ctx, relation.Uid)
		if err != nil {
			return err
		}

		retBalance = account.Balance + in.Amount
		account.Balance = retBalance
		err = txAccount.Update(ctx, account)
		if err != nil {
			return err
		}

		_, err = txFlow.Insert(ctx, &mysql.TAccountFlow{
			Uid:      relation.Uid,
			UserId:   in.UserId,
			FlowId:   flowId,
			FlowType: FlowTypeIn,
			BizType:  BizTypeDeposit,
			Amount:   in.Amount,
		})
		return err
	})
	if err != nil {
		return nil, err
	}

	return &user_mgr_pb.DepositRsp{
		UserId:  in.UserId,
		Balance: retBalance,
	}, nil
}
