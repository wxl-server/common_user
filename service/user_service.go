package service

import (
	"common_user/domain/converter"
	"common_user/repo"
	"context"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/wxl-server/idl_gen/kitex_gen/common_user"
	"go.uber.org/dig"
)

type UserService interface {
	SignUp(ctx context.Context, req *common_user.SignUpReq) (resp *common_user.SignUpResp, err error)
}

type Param struct {
	dig.In
	UserRepo repo.UserRepo
}

type UserServiceImpl struct {
	p Param
}

func NewUserService(p Param) UserService {
	return &UserServiceImpl{
		p: p,
	}
}

func (u UserServiceImpl) SignUp(ctx context.Context, req *common_user.SignUpReq) (resp *common_user.SignUpResp, err error) {
	id, err := u.p.UserRepo.SignUp(ctx, converter.SignUpReqDTO2DO(req))
	if err != nil {
		logger.CtxErrorf(ctx, "SignUp failed, err = %v", err)
		return nil, err
	}
	return &common_user.SignUpResp{
		Id: id,
	}, nil
}
