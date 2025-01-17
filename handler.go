package main

import (
	"context"
	common_user "github.com/qcq1/idl_gen/kitex_gen/common_user"
)

var handler common_user.CommonUser

// CommonUserImpl implements the last service interface defined in the IDL.
type CommonUserImpl struct{}

func NewHandler() common_user.CommonUser {
	return &CommonUserImpl{}
}

// SignUp implements the CommonUserImpl interface.
func (s *CommonUserImpl) SignUp(ctx context.Context, req *common_user.SignUpReq) (resp *common_user.SignUpResp, err error) {
	// TODO: Your code here...
	// logger.CtxInfof(ctx, "req = %v", render.Render(req))
	return
}

// Login implements the CommonUserImpl interface.
func (s *CommonUserImpl) Login(ctx context.Context, req *common_user.LoginReq) (resp *common_user.LoginResp, err error) {
	// TODO: Your code here...
	return &common_user.LoginResp{
		Token: "wxl555555",
	}, nil
}
