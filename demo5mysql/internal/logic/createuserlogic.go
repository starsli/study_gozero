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

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserLogic) CreateUser(in *user_mgr_pb.CreateUserReq) (*user_mgr_pb.CreateUserRsp, error) {
	// 先查询relation是否已经存在
	relation, err := l.svcCtx.TRelationModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		if err != sqlx.ErrNotFound {
			return nil, err
		}
	} else {
		if relation.State == 1 {
			// 继续关联中，不执行后续操作
		} else if relation.State == 2 {
			// 重入关联成功的用户, 要校验关键信息一致性
			userInfoTmp, err := l.svcCtx.TUserInfoModel.FindOne(l.ctx, relation.Uid)
			if err != nil {
				return nil, err
			}

			if userInfoTmp.Name != in.Name || userInfoTmp.Age != int64(in.Age) {
				return nil, errors.New("user info is not consistent")
			}

			return &user_mgr_pb.CreateUserRsp{
				UserId:   in.UserId,
				IsRepeat: 1,
			}, nil
		} else {
			return nil, errors.New("relation state is not 1 or 2")
		}
	}

	// 生成用户ID
	newUid, err := l.generateUid()

	l.Logger.Infof("starsli1 generateUid: %v", newUid)

	if err != nil {
		l.Logger.Errorf("generateUid failed: %v", err)
		return nil, err
	}

	l.Logger.Infof("starsli2 generateUid: %v", newUid)
	l.svcCtx.TRelationModel.Insert(l.ctx, &mysql.TRelation{
		UserId: in.UserId,
		Uid:    newUid,
		State:  1, // 注册中
	})

	l.Logger.Infof("starsli3 generateUid: %v", newUid)
	// 插入用户信息
	_, err = l.svcCtx.TUserInfoModel.Insert(l.ctx, &mysql.TUserInfo{
		Uid:    newUid,
		UserId: in.UserId,
		Name:   in.Name,
		Age:    int64(in.Age),
	})
	if err != nil {
		l.Logger.Errorf("insert user info failed: %v", err)
		return nil, err
	}

	l.Logger.Infof("starsli generateUid: %v", newUid)
	_, err = l.svcCtx.TAccountModel.Insert(l.ctx, &mysql.TAccount{
		Uid:     newUid,
		UserId:  in.UserId,
		Balance: 0,
	})
	if err != nil {
		l.Logger.Errorf("insert account failed: %v", err)
		return nil, err
	}

	l.svcCtx.TRelationModel.Update(l.ctx, &mysql.TRelation{
		UserId: in.UserId,
		Uid:    newUid,
		State:  2, // 注册成功
	})

	return &user_mgr_pb.CreateUserRsp{
		UserId:   in.UserId,
		IsRepeat: 0,
	}, nil
}

// generateUid 从 t_uid_segment 获取一个UID，需要在事务中完成
func (l *CreateUserLogic) generateUid() (int64, error) {
	var newUid int64
	err := l.svcCtx.TUidSegmentModel.TransactCtx(l.ctx, func(ctx context.Context, tx mysql.TUidSegmentModel) error {
		// 查询当前 segment (使用 FOR UPDATE 锁定行)
		segment, err := tx.FindOneForUpdate(ctx, 1)
		if err != nil {
			return err
		}

		// 计算新的 UID
		newUid = segment.UidMax + 1

		// 更新 uid_max
		segment.UidMax = newUid
		return tx.Update(ctx, segment)
	})

	if err != nil {
		return 0, err
	}

	return newUid, nil
}
