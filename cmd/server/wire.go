//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/fxkt-tech/picsal/internal/biz"
	"github.com/fxkt-tech/picsal/internal/conf"
	"github.com/fxkt-tech/picsal/internal/data"
	"github.com/fxkt-tech/picsal/internal/server"
	"github.com/fxkt-tech/picsal/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// initApp init kratos application.
func initApp(*conf.Bootstrap, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		server.ProviderSet,
		data.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
		newApp,
	))
}
