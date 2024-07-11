//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"linkme-post/internal/biz"
	"linkme-post/internal/conf"
	"linkme-post/internal/data"
	"linkme-post/internal/server"
	"linkme-post/internal/service"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *conf.Service, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet1, data.ProviderSet, biz.ProviderSet3, service.ProviderSet4, newApp))
}
