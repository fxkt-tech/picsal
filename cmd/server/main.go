package main

import (
	"os"

	"github.com/fxkt-tech/picsal/internal/conf"
	"github.com/fxkt-tech/picsal/pkg/bootstrap"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
)

var (
	Name    string
	Version string
	id, _   = os.Hostname()
)

func newApp(logger log.Logger, hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(hs),
	)
}

func main() {
	// 初始化各种配置
	bs, err := bootstrap.New[conf.Bootstrap]().
		LoadEnv().
		LoadFlag().
		LoadLogger().
		LoadConfig().
		Result()
	if err != nil {
		panic(err)
	}
	defer bs.Defer()

	logger := bs.Logger()
	config := bs.Config()

	// 运行各种服务
	app, cleanup, err := initApp(&config, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	if err := app.Run(); err != nil {
		panic(err)
	}
}
