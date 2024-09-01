package biz

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-post/events/publish"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"time"
)

const (
	Draft     uint8 = iota // 0: 草稿状态
	Published              // 1: 发布状态
	Withdrawn              // 2: 撤回状态
	Deleted                // 3: 删除状态
)

type Post struct {
	ID           int64        `gorm:"primarykey"`
	CreatedAt    time.Time    `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time    `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt    sql.NullTime `gorm:"index"`
	Title        string       `gorm:"size:255;not null"`            // 文章标题
	Content      string       `gorm:"type:text;not null"`           // 文章内容
	Status       uint8        `gorm:"default:0"`                    // 文章状态，如草稿、发布等
	UserID       int64        `gorm:"column:user_id;index"`         // 用户uid
	Slug         string       `gorm:"size:100;uniqueIndex"`         // 文章的唯一标识，用于生成友好URL
	CategoryID   int64        `gorm:"index"`                        // 关联分类表的外键
	PlateID      int64        `gorm:"index"`                        // 关联板块表的外键
	Plate        Plate        `gorm:"foreignKey:PlateID"`           // 板块关系
	Tags         string       `gorm:"type:varchar(255);default:''"` // 文章标签，以逗号分隔
	CommentCount int64        `gorm:"default:0"`                    // 文章的评论数量
}

type ListPubPost struct {
	ID           int64     `bson:"id"`           // MongoDB的ObjectID
	CreatedAt    time.Time `bson:"createdat"`    // 创建时间
	UpdatedAt    time.Time `bson:"updatedat"`    // 更新时间
	Title        string    `bson:"title"`        // 文章标题
	Content      string    `bson:"content"`      // 文章内容
	Status       uint8     `bson:"status"`       // 文章状态，如草稿、发布等
	UserID       int64     `bson:"userid"`       // 用户uid
	Slug         string    `bson:"slug"`         // 文章的唯一标识，用于生成友好URL
	CategoryID   int64     `bson:"categoryid"`   // 关联分类表的外键
	PlateID      int64     `bson:"plateid"`      // 关联板块表的外键
	Plate        Plate     `bson:"plate"`        // 板块关系
	Tags         string    `bson:"tags"`         // 文章标签，以逗号分隔
	CommentCount int64     `bson:"commentcount"` // 文章的评论数量
}

type Plate struct {
	ID          int64        `gorm:"primaryKey;autoIncrement"`         // 板块ID
	Name        string       `gorm:"size:255;not null;uniqueIndex"`    // 板块名称
	Description string       `gorm:"type:text"`                        // 板块描述
	CreatedAt   time.Time    `gorm:"column:created_at;autoCreateTime"` // 创建时间
	UpdatedAt   time.Time    `gorm:"column:updated_at;autoUpdateTime"` // 更新时间
	DeletedAt   sql.NullTime `gorm:"column:deleted_at;index"`          // 删除时间
	Deleted     bool         `gorm:"column:deleted;default:false"`     // 是否删除
	UserID      int64        `gorm:"column:uid;index"`                 // 板主ID
	Posts       []Post       `gorm:"foreignKey:PlateID"`               // 帖子关系
}

type Pagination struct {
	Page int    // 当前页码
	Size *int64 // 每页数据
	Uid  int64
	// 以下字段通常在服务端内部使用，不需要客户端传递
	Offset *int64 // 数据偏移量
	Total  *int64 // 总数据量
}

type PostData interface {
	Insert(ctx context.Context, post Post) (int64, error)                      // 创建一个新的帖子记录
	UpdateById(ctx context.Context, post Post) error                           // 根据ID更新一个帖子记录
	UpdateStatus(ctx context.Context, post Post) error                         // 更新帖子的状态
	GetById(ctx context.Context, postId int64, uid int64) (Post, error)        // 根据ID获取一个帖子记录
	GetPubById(ctx context.Context, postId int64) (ListPubPost, error)         // 根据ID获取一个已发布的帖子记录
	ListPub(ctx context.Context, pagination Pagination) ([]ListPubPost, error) // 获取已发布的帖子记录列表
	List(ctx context.Context, pagination Pagination) ([]Post, error)           // 获取个人的帖子记录列表
	DeleteById(ctx context.Context, post Post) error
	CreatePlate(ctx context.Context, plate Plate) error
	ListPlates(ctx context.Context, pagination Pagination) ([]Plate, error)
	UpdatePlate(ctx context.Context, plate Plate) error
	DeletePlate(ctx context.Context, plateId int64) error
}

type PostBiz struct {
	postData        PostData
	l               *zap.Logger
	publishProducer publish.Producer
}

func NewPostBiz(postData PostData, l *zap.Logger, publishProducer publish.Producer) *PostBiz {
	return &PostBiz{
		postData:        postData,
		publishProducer: publishProducer,
		l:               l,
	}
}

func (pb *PostBiz) CreatePost(ctx context.Context, post Post) (int64, error) {
	post.Slug = uuid.New().String()
	post.Status = Draft // 默认状态为草稿

	return pb.postData.Insert(ctx, post)
}

func (pb *PostBiz) UpdatePost(ctx context.Context, post Post) error {
	post.Status = Draft // 默认状态为草稿
	return pb.postData.UpdateById(ctx, post)
}

func (pb *PostBiz) UpdatePostStatus(ctx context.Context, post Post) error {
	return pb.postData.UpdateStatus(ctx, post)
}

func (pb *PostBiz) DeletePost(ctx context.Context, post Post) error {
	return pb.postData.DeleteById(ctx, post)
}

func (pb *PostBiz) GetPost(ctx context.Context, postId int64, uid int64) (Post, error) {
	return pb.postData.GetById(ctx, postId, uid)
}

func (pb *PostBiz) GetPubPost(ctx context.Context, postId int64) (ListPubPost, error) {
	return pb.postData.GetPubById(ctx, postId)
}

func (pb *PostBiz) ListPost(ctx context.Context, pagination Pagination) ([]Post, error) {
	offset := int64(pagination.Page-1) * *pagination.Size
	pagination.Offset = &offset
	return pb.postData.List(ctx, pagination)
}

func (pb *PostBiz) ListPubPost(ctx context.Context, pagination Pagination) ([]ListPubPost, error) {
	offset := int64(pagination.Page-1) * *pagination.Size
	pagination.Offset = &offset
	return pb.postData.ListPub(ctx, pagination)
}

func (pb *PostBiz) PublishPost(ctx context.Context, post Post) error {
	pd, err := pb.postData.GetById(ctx, post.ID, post.UserID)
	if err != nil {
		return fmt.Errorf("get post failed: %w", err)
	}

	// 使用 context.WithCancel 来管理 goroutine 的生命周期
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// 使用装饰器封装 goroutine 逻辑
	asyncPublish := pb.withAsyncCancel(ctx, cancel, func() error {
		if err := pb.publishProducer.ProducePublishEvent(publish.PublishEvent{
			PostId:  pd.ID,
			Content: pd.Content,
			Title:   pd.Title,
			UserID:  pd.UserID,
		}); err != nil {
			return fmt.Errorf("failed to produce publish event: %w", err)
		}
		return nil
	})

	asyncPublish()

	return nil
}

func (pb *PostBiz) ListPlates(ctx context.Context, pagination Pagination) ([]Plate, error) {
	offset := int64(pagination.Page-1) * *pagination.Size
	pagination.Offset = &offset
	return pb.postData.ListPlates(ctx, pagination)
}

func (pb *PostBiz) CreatePlate(ctx context.Context, plate Plate) error {
	return pb.postData.CreatePlate(ctx, plate)
}

func (pb *PostBiz) UpdatePlate(ctx context.Context, plate Plate) error {
	return pb.postData.UpdatePlate(ctx, plate)
}

func (pb *PostBiz) DeletePlate(ctx context.Context, plateId int64) error {
	return pb.postData.DeletePlate(ctx, plateId)
}

// withAsyncCancel 装饰器函数，用来封装 goroutine 逻辑并处理错误和取消操作
func (pb *PostBiz) withAsyncCancel(_ context.Context, cancel context.CancelFunc, fn func() error) func() {
	return func() {
		go func() {
			// 确保 goroutine 中的 panic 不会导致程序崩溃
			defer func() {
				if r := recover(); r != nil {
					pb.l.Error("panic occurred in async operation goroutine", zap.Any("error", r))
					cancel() // 取消操作
				}
			}()

			// 执行目标函数
			if err := fn(); err != nil {
				pb.l.Error("async operation failed", zap.Error(err))
				cancel() // 取消操作
			}
		}()
	}
}
