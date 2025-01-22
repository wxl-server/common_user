package repo

import (
	"common_user/domain"
	"common_user/domain/converter"
	"common_user/sal/dao/generator/query"
	"context"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/wxl-server/common/id_gen"

	"go.uber.org/dig"
)

type UserRepo interface {
	CountUser(ctx context.Context, email string) (count int64, err error)
	CreateUser(ctx context.Context, do *domain.UserDO) (id int64, err error)
	QueryUser(ctx context.Context, email string) (do *domain.UserDO, err error)
	UpdatePassword(ctx context.Context, do *domain.UserDO) (err error)
}

type Param struct {
	dig.In
}

type UserRepoImpl struct {
	p Param
}

func NewUserRepo(p Param) UserRepo {
	return &UserRepoImpl{
		p: p,
	}
}

func (u UserRepoImpl) UpdatePassword(ctx context.Context, do *domain.UserDO) (err error) {
	po := query.Q.UserPO
	_, err = po.WithContext(ctx).Where(po.ID.Eq(do.ID)).UpdateSimple(po.Password.Value(do.Password))
	if err != nil {
		logger.CtxErrorf(ctx, "UpdatePassword failed, err = %v", err)
		return err
	}
	return nil
}

func (u UserRepoImpl) QueryUser(ctx context.Context, email string) (do *domain.UserDO, err error) {
	po := query.Q.UserPO
	userPO, err := po.WithContext(ctx).Where(po.Email.Eq(email)).First()
	if err != nil {
		logger.CtxErrorf(ctx, "QueryUser failed, err = %v", err)
		return nil, err
	}
	return converter.UserPO2DO(userPO), nil
}

func (u UserRepoImpl) CountUser(ctx context.Context, email string) (count int64, err error) {
	po := query.Q.UserPO
	count, err = po.WithContext(ctx).Where(po.Email.Eq(email)).Count()
	if err != nil {
		logger.CtxErrorf(ctx, "CountUser failed, err = %v", err)
		return 0, err
	}
	return count, nil
}

func (u UserRepoImpl) CreateUser(ctx context.Context, do *domain.UserDO) (id int64, err error) {
	po := query.Q.UserPO
	do.ID, err = id_gen.NextID()
	if err != nil {
		logger.CtxErrorf(ctx, "generate id failed, err = %v", err)
		return 0, err
	}
	err = po.WithContext(ctx).Create(converter.UserDO2PO(do))
	if err != nil {
		logger.CtxErrorf(ctx, "CreateUser failed, err = %v", err)
		return 0, err
	}
	return do.ID, nil
}
