package logic

import (
	"context"
	"errors"

	"demo5mysql/internal/svc"
	"demo5mysql/user_mgr_pb"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type GetUserFlowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserFlowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserFlowLogic {
	return &GetUserFlowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserFlowLogic) GetUserFlow(in *user_mgr_pb.GetUserFlowReq) (*user_mgr_pb.GetUserFlowRsp, error) {
	// 校验字段长度
	// UserId[1,64]
	if len(in.UserId) < 1 || len(in.UserId) > 64 {
		return nil, errors.New("userId length is not in range [1,64]")
	}

	// Limit[0,100]
	if in.Limit > 100 || in.Limit < 0 {
		return nil, errors.New("limit is not in range [0,100]")
	}

	// 没有限制，默认查询10条数据
	if in.Limit <= 0 {
		in.Limit = 10
	}

	// 没有指定偏移量，默认从0开始查询
	if in.Offset < 0 {
		in.Offset = 0
	}

	// 校验用户是否存在
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

	// 查询用户流水, 根据create_time排序, 从新到旧查询
	flows, err := l.svcCtx.TAccountFlowModel.FindByUidWithOrderByTime(l.ctx, relation.Uid, in.Offset, in.Limit)
	if err != nil {
		return nil, err
	}

	flowItems := make([]*user_mgr_pb.GetUserFlowItemRsp, 0, len(flows))
	for _, flow := range flows {
		flowItems = append(flowItems, &user_mgr_pb.GetUserFlowItemRsp{
			FlowId:     flow.FlowId,
			UserId:     flow.UserId,
			Amount:     flow.Amount,
			BizType:    int32(flow.BizType),
			FlowType:   int32(flow.FlowType),
			CreateTime: flow.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	// 判断是否还有更多数据
	flag := int32(0)
	if int32(len(flows)) < in.Limit {
		flag = 2
	} else {
		flag = 1
	}

	return &user_mgr_pb.GetUserFlowRsp{
		UserId: in.UserId,
		Flow:   flowItems,
		Flag:   flag,                          // 是否还有更多数据 1:有 2:无
		Offset: in.Offset + int32(len(flows)), // 下次查询的偏移量
	}, nil
}
