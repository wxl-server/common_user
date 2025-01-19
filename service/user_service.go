package service

import (
	"common_user/biz_error"
	"common_user/domain/converter"
	"common_user/repo"
	"context"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/wxl-server/idl_gen/kitex_gen/common_user"
	"go.uber.org/dig"
)

type UserService interface {
	SignUp(ctx context.Context, req *common_user.SignUpReq) (resp *common_user.SignUpResp, err error)
	UpdatePassword(ctx context.Context, req *common_user.UpdatePasswordReq) (*common_user.UpdatePasswordResp, error)
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
	count, err := u.p.UserRepo.CountUser(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		logger.CtxErrorf(ctx, "sign up failed, email has been used, email = %s", req.Email)
		return nil, biz_error.SignUpError
	}
	id, err := u.p.UserRepo.CreateUser(ctx, converter.SignUpReqDTO2DO(req))
	if err != nil {
		logger.CtxErrorf(ctx, "CreateUser failed, err = %v", err)
		return nil, err
	}
	return &common_user.SignUpResp{
		Id: id,
	}, nil
}

func (u UserServiceImpl) UpdatePassword(ctx context.Context, req *common_user.UpdatePasswordReq) (*common_user.UpdatePasswordResp, error) {
	count, err := u.p.UserRepo.QueryUser(ctx, req.Email)
}
