//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-interactive/internal/biz"
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-interactive/internal/conf"
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-interactive/internal/data"
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-interactive/internal/server"
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-interactive/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *conf.Service, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
