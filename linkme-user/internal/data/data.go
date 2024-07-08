package data

import (
	"fmt"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	clientv3 "go.etcd.io/etcd/client/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"linkme-user/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewUserRepo, NewJWT, NewRedis, NewLogger, NewSnowflake, NewEtcd)

// Data .
type Data struct {
	db    *gorm.DB
	redis redis.Cmdable
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, db *gorm.DB) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
		sqlDB, err := db.DB()
		if err != nil {
			log.NewHelper(logger).Error("failed to get sql.DB from gorm.DB", err)
		}
		if er := sqlDB.Close(); er != nil {
			log.NewHelper(logger).Error("failed to close database connection", er)
		}
	}
	return &Data{db: db}, cleanup, nil
}

// NewDB 初始化数据库
func NewDB(c *conf.Data) (*gorm.DB, error) {
	dbConfig := c.Database
	db, err := gorm.Open(mysql.Open(dbConfig.Source), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}
	// 初始化表
	if err = InitTables(db); err != nil {
		return nil, fmt.Errorf("failed to initialize tables: %w", err)
	}
	return db, nil
}

// NewRedis 初始化Redis
func NewRedis(c *conf.Data) redis.Cmdable {
	return redis.NewClient(&redis.Options{
		Addr: c.Redis.Addr,
	})
}

// NewEtcd 初始化Etcd
func NewEtcd(c *conf.Data, logger log.Logger, gs *grpc.Server, hs *http.Server) registry.Registrar {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   c.Etcd.Endpoints,
		DialTimeout: c.Etcd.Timeout.AsDuration(),
	})
	if err != nil {
		panic(err)
	}
	reg := etcd.New(client)
	kratos.New(kratos.Name("linkme-user"),
		kratos.Version("v1.0"),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(hs, gs),
		kratos.Registrar(reg),
	)
	return etcd.New(client)
}
