package service

import (
	"common_user/biz_error"
	"common_user/domain/converter"
	"common_user/repo"
	"common_user/sal/jwt"
	"context"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/wxl-server/idl_gen/kitex_gen/common_user"
	"go.uber.org/dig"
)

type UserService interface {
	SignUp(ctx context.Context, req *common_user.SignUpReq) (resp *common_user.SignUpResp, err error)
	UpdatePassword(ctx context.Context, req *common_user.UpdatePasswordReq) (*common_user.UpdatePasswordResp, error)
	Login(ctx context.Context, req *common_user.LoginReq) (resp *common_user.LoginResp, err error)
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
	if req.OldPassword == req.Password {
		logger.CtxErrorf(ctx, "UpdatePassword failed, err = %v", biz_error.UpdatePasswordError)
		return nil, biz_error.UpdatePasswordError
	}
	do, err := u.p.UserRepo.QueryUser(ctx, req.Email)
	if err != nil {
		logger.CtxErrorf(ctx, "QueryUser failed, err = %v", err)
		return nil, err
	}
	if do.Password != req.OldPassword {
		logger.CtxErrorf(ctx, "UpdatePassword failed, err = %v", biz_error.UpdatePasswordError2)
		return nil, biz_error.UpdatePasswordError2
	}
	do.Password = req.Password
	err = u.p.UserRepo.UpdatePassword(ctx, do)
	if err != nil {
		logger.CtxErrorf(ctx, "UpdatePassword failed, err = %v", err)
		return nil, err
	}
	return &common_user.UpdatePasswordResp{}, nil
}

func (u UserServiceImpl) Login(ctx context.Context, req *common_user.LoginReq) (resp *common_user.LoginResp, err error) {
	do, err := u.p.UserRepo.QueryUser(ctx, req.Email)
	if err != nil {
		logger.CtxErrorf(ctx, "QueryUser failed, err = %v", err)
		return nil, biz_error.LoginError
	}
	if do.Password != req.Password {
		logger.CtxErrorf(ctx, "login failed, password is wrong")
		return nil, biz_error.LoginError2
	}
	token, err := jwt.GenerateToken(ctx, do.ID)
	if err != nil {
		logger.CtxErrorf(ctx, "login failed, err = %v", err)
		return nil, biz_error.LoginError
	}
	return &common_user.LoginResp{
		Token: token,
	}, nil
}
