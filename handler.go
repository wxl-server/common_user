package main

import (
	"common_user/service"
	"context"

	common_user "github.com/wxl-server/idl_gen/kitex_gen/common_user"
	"go.uber.org/dig"
)

var handler common_user.CommonUser

// Handler implements the last service interface defined in the IDL.
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

// SignUp implements the Handler interface.
func (s *Handler) SignUp(ctx context.Context, req *common_user.SignUpReq) (resp *common_user.SignUpResp, err error) {
	return s.p.UserService.SignUp(ctx, req)
}

// Login implements the Handler interface.
func (s *Handler) Login(ctx context.Context, req *common_user.LoginReq) (resp *common_user.LoginResp, err error) {
	// TODO: Your code here...
	return &common_user.LoginResp{
		Token: "wxl555555",
	}, nil
}
