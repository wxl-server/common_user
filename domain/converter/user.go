package converter

import (
	"common_user/domain"

	"github.com/wxl-server/idl_gen/kitex_gen/common_user"
)

func SignUpReqDTO2DO(req *common_user.SignUpReq) *domain.SignUpReqDO {
	return &domain.SignUpReqDO{
		Email:    req.Email,
		Password: req.Password,
	}
}
