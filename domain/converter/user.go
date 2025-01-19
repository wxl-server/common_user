package converter

import (
	"common_user/domain"
	"common_user/sal/dao/generator/model"

	"github.com/wxl-server/idl_gen/kitex_gen/common_user"
)

func SignUpReqDTO2DO(req *common_user.SignUpReq) *domain.UserDO {
	return &domain.UserDO{
		Email:    req.Email,
		Password: req.Password,
	}
}

func UserDO2PO(do *domain.UserDO) *model.UserPO {
	return &model.UserPO{
		ID:        do.ID,
		Email:     do.Email,
		Password:  do.Password,
		Extra:     do.Extra,
		CreatedAt: do.CreatedAt,
		UpdatedAt: do.UpdatedAt,
	}
}
