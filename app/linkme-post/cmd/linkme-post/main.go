package main

import (
	"context"
	"flag"
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-post/events/publish"
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-post/events/sync"
	"os"
	"os/signal"
	"syscall"

	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.15.0"

	checkpb "github.com/GoSimplicity/LinkMe-microservices/api/check/v1"
	userpb "github.com/GoSimplicity/LinkMe-microservices/api/user/v1"
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-post/internal/conf"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/uuid"
	clientv3 "go.etcd.io/etcd/client/v3"

	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name = "linkme-post"
	// Version is the version of the compiled software.
	Version = "v1.0.0"
	// flagconf is the config flag.
	flagconf string
	// id, _ = os.Hostname()
	id = Name + "-" + uuid.NewString()
)

func init() {
	//flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
	flag.StringVar(&flagconf, "conf", "./configs", "config path, eg: -conf config.yaml")
}

func newApp(cs *conf.Service, logger log.Logger, gs *grpc.Server, hs *http.Server, publishConsumer *publish.PublishPostEventConsumer, syncConsumer *sync.SyncConsumer) *kratos.App {
	// 在New前完成初始化调用
	reg := initServiceRegistry(cs)

	// 启动 Kafka 消费者
	go func() {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		publishConsumer.Start(ctx)

		sigterm := make(chan os.Signal, 1)
		signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
		<-sigterm
		cancel()
	}()

	go func() {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		syncConsumer.Start(ctx)

		sigterm := make(chan os.Signal, 1)
		signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
		<-sigterm
		cancel()
	}()

	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),
		kratos.Registrar(reg),
	)
}

func main() {
	flag.Parse()
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	// 设置 Jaeger 追踪导出器，用于收集追踪信息
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(bc.Trace.Endpoint)))
	if err != nil {
		panic(err)
	}

	// 创建 OpenTelemetry 追踪提供者
	tp := tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(exp), // 批量处理追踪数据
		tracesdk.WithResource(resource.NewSchemaless(
			semconv.ServiceNameKey.String(Name), // 设置服务名称
		)),
	)

	userClient, err := initUserClient(bc.Service, logger)
	if err != nil {
		panic(err)
	}

	checkClient, err := initCheckClient(bc.Service, logger)
	if err != nil {
		panic(err)
	}

	app, cleanup, err := wireApp(bc.Server, bc.Data, bc.Service, userClient, checkClient, logger, tp)
	if err != nil {
		panic(err)
	}

	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}

// 初始化服务注册
func initServiceRegistry(c *conf.Service) registry.Registrar {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   c.Etcd.Addr,
		DialTimeout: c.Etcd.Timeout.AsDuration(),
	})
	if err != nil {
		panic(err)
	}
	reg := etcd.New(client)
	return reg
}

// 初始化用户客户端
func initUserClient(c *conf.Service, logger log.Logger) (userpb.UserClient, error) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   c.Etcd.Addr,
		DialTimeout: c.Etcd.Timeout.AsDuration(),
	})
	if err != nil {
		return nil, err
	}
	r := etcd.New(client)
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///linkme-user"),
		grpc.WithDiscovery(r),
		grpc.WithLogger(logger),
	)
	if err != nil {
		return nil, err
	}
	return userpb.NewUserClient(conn), nil
}

// 初始化审核客户端
func initCheckClient(c *conf.Service, logger log.Logger) (checkpb.CheckClient, error) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   c.Etcd.Addr,
		DialTimeout: c.Etcd.Timeout.AsDuration(),
	})
	if err != nil {
		return nil, err
	}
	r := etcd.New(client)
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///linkme-check"),
		grpc.WithDiscovery(r),
		grpc.WithLogger(logger),
	)
	if err != nil {
		return nil, err
	}
	return checkpb.NewCheckClient(conn), nil
}
