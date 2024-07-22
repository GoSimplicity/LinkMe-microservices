//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	pb "github.com/GoSimplicity/LinkMe-microservices/api/post/v1"
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-check/internal/biz"
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-check/internal/conf"
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-check/internal/data"
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-check/internal/server"
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-check/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *conf.Service, pb.PostClient, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}