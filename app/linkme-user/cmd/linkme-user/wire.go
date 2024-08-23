//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-user/internal/biz"
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-user/internal/conf"
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-user/internal/data"
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-user/internal/middleware"
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-user/internal/server"
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-user/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *conf.Service, log.Logger, *tracesdk.TracerProvider) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, middleware.ProviderSet, newApp))
}
