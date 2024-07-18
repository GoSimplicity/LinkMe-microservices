package data

import (
	"context"
	"fmt"
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-post/internal/conf"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewRedis, NewLogger, NewMongoDB, NewPostData)

// Data .
type Data struct {
	db    *gorm.DB
	redis redis.Cmdable
	mongo *mongo.Client
	cs    *conf.Service
}

// NewData .
func NewData(c *conf.Data, cs *conf.Service, db *gorm.DB, redis redis.Cmdable, logger log.Logger, mongo *mongo.Client) (*Data, func(), error) {
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
	return &Data{db: db, redis: redis, cs: cs, mongo: mongo}, cleanup, nil
}

// NewDB initializes the database.
func NewDB(c *conf.Data) (*gorm.DB, error) {
	dbConfig := c.Database
	db, err := gorm.Open(mysql.Open(dbConfig.Source), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}
	// Initialize tables
	if err = InitTables(db); err != nil {
		return nil, fmt.Errorf("failed to initialize tables: %w", err)
	}
	return db, nil
}

// NewRedis initializes Redis.
func NewRedis(c *conf.Data) redis.Cmdable {
	return redis.NewClient(&redis.Options{
		Addr: c.Redis.Addr,
	})
}

func NewMongoDB(c *conf.Data) *mongo.Client {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(c.Mongo.Addr))
	if err != nil {
		panic(err)
	}
	return client
}
