package logic

import (
	"context"

	"demo5mysql/internal/svc"
	"demo5mysql/user_mgr_pb"

	"github.com/zeromicro/go-zero/core/logx"
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
	if err := CheckUserId(in.UserId); err != nil {
		return nil, err
	}

	limit, err := CheckLimit(in.Limit)
	if err != nil {
		return nil, err
	}
	offset := CheckOffset(in.Offset)

	relation, err := CheckUserRegistered(l.ctx, l.svcCtx.TRelationModel, in.UserId)
	if err != nil {
		return nil, err
	}

	// 查询用户流水, 根据create_time排序, 从新到旧查询
	flows, err := l.svcCtx.TAccountFlowModel.FindByUidWithOrderByTime(l.ctx, relation.Uid, offset, limit)
	if err != nil {
		return nil, err
	}

	flowItems := make([]*user_mgr_pb.GetUserFlowItemRsp, 0, len(flows))
	for _, flow := range flows {
		flowItems = append(flowItems, &user_mgr_pb.GetUserFlowItemRsp{
			FlowId:         flow.FlowId,
			UserId:         flow.UserId,
			Amount:         flow.Amount,
			BizType:        int32(flow.BizType),
			FlowType:       int32(flow.FlowType),
			CreateTime:     flow.CreateTime.Format("2006-01-02 15:04:05"),
			CounterpartyId: flow.CounterpartyId,
		})
	}

	// 判断是否还有更多数据
	flag := int32(0)
	if int32(len(flows)) < limit {
		flag = 2
	} else {
		flag = 1
	}

	return &user_mgr_pb.GetUserFlowRsp{
		UserId: in.UserId,
		Flow:   flowItems,
		Flag:   flag,                       // 是否还有更多数据 1:有 2:无
		Offset: offset + int32(len(flows)), // 下次查询的偏移量
	}, nil
}
