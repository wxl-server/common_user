package main

import (
	"common_user/service"
	"context"

	common_user "github.com/wxl-server/idl_gen/kitex_gen/common_user"
	"go.uber.org/dig"
)

var handler common_user.CommonUser

type Handler struct {
	p Param
}

type Param struct {
	dig.In
	UserService service.UserService
}

func NewHandler(p Param) {
	handler = &Handler{
		p: p,
	}
}

func (s *Handler) SignUp(ctx context.Context, req *common_user.SignUpReq) (resp *common_user.SignUpResp, err error) {
	return s.p.UserService.SignUp(ctx, req)
}

func (s *Handler) UpdatePassword(ctx context.Context, req *common_user.UpdatePasswordReq) (resp *common_user.UpdatePasswordResp, err error) {
	return s.p.UserService.UpdatePassword(ctx, req)
}

func (s *Handler) Login(ctx context.Context, req *common_user.LoginReq) (resp *common_user.LoginResp, err error) {
	return s.p.UserService.Login(ctx, req)
}

func (s *Handler) ValidateToken(ctx context.Context, req *common_user.ValidateTokenReq) (resp *common_user.ValidateTokenResp, err error) {
	return s.p.UserService.ValidateToken(ctx, req)
}
