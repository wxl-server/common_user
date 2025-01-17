package main

import (
	"github.com/qcq1/common/wxl_cluster"
	common_user "github.com/qcq1/idl_gen/kitex_gen/common_user/commonuser"
)

func main() {
	wxl_cluster.NewServer(common_user.NewServer, NewHandler(), "common_user", 8090)
}
