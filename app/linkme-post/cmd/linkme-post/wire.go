//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	checkpb "github.com/GoSimplicity/LinkMe-microservices/api/check/v1"
	userpb "github.com/GoSimplicity/LinkMe-microservices/api/user/v1"
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-post/events"
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-post/internal/biz"
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-post/internal/conf"
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-post/internal/data"
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-post/internal/server"
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-post/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *conf.Service, userpb.UserClient, checkpb.CheckClient, log.Logger, *tracesdk.TracerProvider) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, events.ProviderSet, newApp))
}
