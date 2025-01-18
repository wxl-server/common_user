package main

import (
	"common_user/sal/dao"
	"common_user/service"
	"context"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/wxl-server/common/render"
	"github.com/wxl-server/common/wxl_cluster"
	common_user "github.com/wxl-server/idl_gen/kitex_gen/common_user/commonuser"
	"go.uber.org/dig"
)

var (
	initCtx   = context.Background()
	container = dig.New()
)

func main() {
	wxl_cluster.NewServer(common_user.NewServer, handler, "common_user", 8090)
}

func InitContainer() {
	// context
	{
		mustProvide(func() context.Context { return initCtx })
	}

	// db
	{
		mustInvoke(dao.InitDB)
	}

	// service
	{
		mustProvide(service.NewUserService)
	}

	// handler
	{
		mustInvoke(NewHandler)
	}
}

func mustProvide(constructor interface{}, opts ...dig.ProvideOption) {
	if err := container.Provide(constructor, opts...); err != nil {
		logger.Errorf("container provide failed, err = %v, constructor = %v", err, render.Render(constructor))
		panic(err)
	}
}

func mustInvoke(function interface{}, opts ...dig.InvokeOption) {
	if err := container.Invoke(function, opts...); err != nil {
		logger.Errorf("container invoke failed, err = %v, function = %v", err, render.Render(function))
		panic(err)
	}
}
