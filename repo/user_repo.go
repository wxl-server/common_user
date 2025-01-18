package repo

import (
	"common_user/domain"
	"common_user/sal/dao/generator/query"
	"context"
	"github.com/bytedance/gopkg/util/logger"

	"go.uber.org/dig"
)

type UserRepo interface {
	CountUser(ctx context.Context, email string) (count int64, err error)
	SignUp(ctx context.Context, do *domain.SignUpReqDO) (id int64, err error)
}

type Param struct {
	dig.In
}

type UserRepoImpl struct {
	p Param
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

func NewUserService(p Param) UserRepo {
	return &UserRepoImpl{
		p: p,
	}
}

func (u UserRepoImpl) SignUp(ctx context.Context, do *domain.SignUpReqDO) (id int64, err error) {
	po := query.Q.UserPO
	po.WithContext(ctx).FirstOrCreate()
}
