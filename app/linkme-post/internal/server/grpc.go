package server

import (
	v1 "github.com/GoSimplicity/LinkMe-microservices/api/post/v1"
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-post/internal/conf"
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-post/internal/service"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, sp *service.PostService, logger log.Logger, tp *tracesdk.TracerProvider) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			// 恢复中间件：用于捕获服务器中的 panic 并防止服务崩溃
			recovery.Recovery(),
			// 追踪中间件：用于分布式追踪，集成 OpenTelemetry
			tracing.Server(
				tracing.WithTracerProvider(tp)),
			// 日志中间件：记录请求的日志信息
			logging.Server(logger),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterPostServer(srv, sp)
	return srv
}
