package data

import (
	"context"
	"time"

	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-interactive/domain"
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-interactive/internal/biz"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type InteractiveData struct {
	db *gorm.DB
	l  *zap.Logger
}

type Interactive struct {
	Id           int64  `gorm:"column:id;primaryKey;autoIncrement"`
	BizId        int64  `gorm:"column:biz_id"`
	BizName      string `gorm:"column:biz_name;type:varchar(255)"`
	ReadCount    int64  `gorm:"column:read_count"`
	LikeCount    int64  `gorm:"column:like_count"`
	CollectCount int64  `gorm:"column:collect_count"`
	UpdateTime   int64  `gorm:"column:update_time"`
	CreateTime   int64  `gorm:"column:create_time"`
	PostId       int64  `gorm:"column:post_id"`
	DeletedAt    int64  `gorm:"index"`
}

// UserLike 用户点赞
type UserLike struct {
	ID         int64  `gorm:"primaryKey;autoIncrement"`                     // 点赞记录ID，主键，自增
	Uid        int64  `gorm:"index"`                                        // 用户ID，用于标识哪个用户点赞
	BizID      int64  `gorm:"index"`                                        // 业务ID，用于标识点赞的业务对象
	BizName    string `gorm:"type:varchar(255)"`                            // 业务名称
	Status     int    `gorm:"type:int"`                                     // 状态，用于表示点赞的状态（如有效、无效等）
	UpdateTime int64  `gorm:"column:updated_at;type:bigint;not null;index"` // 更新时间，Unix时间戳
	CreateTime int64  `gorm:"column:created_at;type:bigint"`                // 创建时间，Unix时间戳
	Deleted    bool   `gorm:"column:deleted;default:false"`                 // 删除标志，表示该记录是否被删除
}

// UserCollection 用户收藏
type UserCollection struct {
	ID           int64  `gorm:"primaryKey;autoIncrement"`                     // 收藏记录ID，主键，自增
	Uid          int64  `gorm:"index"`                                        // 用户ID，用于标识哪个用户收藏
	BizID        int64  `gorm:"index"`                                        // 业务ID，用于标识收藏的业务对象
	BizName      string `gorm:"type:varchar(255)"`                            // 业务名称
	Status       int    `gorm:"column:status"`                                // 状态，用于表示收藏的状态（如有效、无效等）
	CollectionId int64  `gorm:"index"`                                        // 收藏ID，用于标识具体的收藏对象
	UpdateTime   int64  `gorm:"column:updated_at;type:bigint;not null;index"` // 更新时间，Unix时间戳
	CreateTime   int64  `gorm:"column:created_at;type:bigint"`                // 创建时间，Unix时间戳
	Deleted      bool   `gorm:"column:deleted;default:false"`                 // 删除标志，表示该记录是否被删除
}

func NewInteractiveData(db *gorm.DB, l *zap.Logger) biz.InteractiveRepo {
	return &InteractiveData{
		db: db,
		l:  l,
	}
}

func (i *InteractiveData) getCurrentTime() int64 {
	return time.Now().UnixMilli()
}

// AddCollectCount implements biz.InteractiveRepo.
func (i *InteractiveData) AddCollectCount(ctx context.Context, postId int64, biz string) error {
	now := i.getCurrentTime()
	// 创建Interactive实例，用于存储收藏计数更新
	interactive := Interactive{
		BizName:      "post",
		CollectCount: 1,
		PostId:       postId,
		CreateTime:   now,
		UpdateTime:   now,
	}
	// 使用Clauses处理数据库冲突，即在记录已存在时更新收藏计数
	return i.db.WithContext(ctx).Clauses(clause.OnConflict{
		DoUpdates: clause.Assignments(map[string]interface{}{
			"collect_count": gorm.Expr("collect_count + 1"),
			"updated_at":    now,
		}),
	}).Create(&interactive).Error
}

// AddLikeCount implements biz.InteractiveRepo.
func (i *InteractiveData) AddLikeCount(ctx context.Context, postId int64, biz string) error {
	now := i.getCurrentTime()
	// 创建Interactive实例，用于存储点赞计数更新
	interactive := Interactive{
		BizName:    "post",
		LikeCount:  1,
		PostId:     postId,
		CreateTime: now,
		UpdateTime: now,
	}
	// 使用Clauses处理数据库冲突，即在记录已存在时更新点赞计数
	return i.db.WithContext(ctx).Clauses(clause.OnConflict{
		DoUpdates: clause.Assignments(map[string]interface{}{
			"like_count": gorm.Expr("like_count + 1"),
			"updated_at": now,
		}),
	}).Create(&interactive).Error
}

// AddReadCount implements biz.InteractiveRepo.
func (i *InteractiveData) AddReadCount(ctx context.Context, postId int64, biz string) error {
	now := i.getCurrentTime()
	// 创建Interactive实例，用于存储阅读计数更新
	interactive := Interactive{
		BizName:    "post",
		ReadCount:  1,
		PostId:     postId,
		CreateTime: now,
		UpdateTime: now,
	}
	// 使用Clauses处理数据库冲突，即在记录已存在时更新阅读计数
	return i.db.WithContext(ctx).Clauses(clause.OnConflict{
		DoUpdates: clause.Assignments(map[string]interface{}{
			"read_count": gorm.Expr("read_count + 1"),
			"updated_at": now,
		}),
	}).Create(&interactive).Error
}

// GetInteractive implements biz.InteractiveRepo.
func (i *InteractiveData) GetInteractive(ctx context.Context, postId int64) (domain.Interactive, error) {
	var interactive domain.Interactive
	err := i.db.WithContext(ctx).Where("post_id = ?", postId).First(&interactive).Error
	if err != nil {
		return domain.Interactive{}, err
	}
	return interactive, nil
}

// ListInteractive implements biz.InteractiveRepo.
func (i *InteractiveData) ListInteractive(ctx context.Context, pagination domain.Pagination) ([]domain.Interactive, error) {
	var interactives []domain.Interactive
	intSize := int(*pagination.Size)
	intOffset := int(*pagination.Offset)
	err := i.db.WithContext(ctx).Limit(intSize).Offset(intOffset).Find(&interactives).Error
	if err != nil {
		return nil, err
	}
	return interactives, nil
}
