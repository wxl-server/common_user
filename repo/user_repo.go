package repo

import (
	"common_user/domain"
	"context"

	"go.uber.org/dig"
)

type UserRepo interface {
	SignUp(ctx context.Context, do *domain.SignUpReqDO) (id int64, err error)
}

type Param struct {
	dig.In
}

type UserRepoImpl struct {
	p Param
}

func NewUserService(p Param) UserRepo {
	return &UserRepoImpl{
		p: p,
	}
}

func (u UserRepoImpl) SignUp(ctx context.Context, do *domain.SignUpReqDO) (id int64, err error) {
	//TODO implement me
	panic("implement me")
}
