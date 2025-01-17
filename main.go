package main

import (
	"github.com/wxl-server/common/wxl_cluster"
	common_user "github.com/wxl-server/idl_gen/kitex_gen/common_user/commonuser"
)

func main() {
	wxl_cluster.NewServer(common_user.NewServer, NewHandler(), "common_user", 8090)
}
