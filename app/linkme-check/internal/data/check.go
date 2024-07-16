package data

import (
	"context"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"linkme-check/domain"
	"linkme-check/internal/biz"
)

type Check struct {
	ID        int64  `gorm:"primaryKey;autoIncrement"`                     // 审核ID
	PostID    int64  `gorm:"not null"`                                     // 帖子ID
	Content   string `gorm:"type:text;not null"`                           // 审核内容
	Title     string `gorm:"size:255;not null"`                            // 审核标签
	Author    int64  `gorm:"column:author_id;index"`                       // 提交审核的用户ID
	Status    string `gorm:"size:20;not null;default:'Pending'"`           // 审核状态
	Remark    string `gorm:"type:text"`                                    // 审核备注
	CreatedAt int64  `gorm:"column:created_at;type:bigint;not null"`       // 创建时间
	UpdatedAt int64  `gorm:"column:updated_at;type:bigint;not null;index"` // 更新时间
}

type checkData struct {
	db *gorm.DB
	l  *zap.Logger
}

func NewCheckData(db *gorm.DB, l *zap.Logger) biz.CheckData {
	return &checkData{
		db: db,
		l:  l,
	}
}

func (c checkData) CreateCheck(ctx context.Context, check domain.Check) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (c checkData) DeleteCheck(ctx context.Context, checkId int64) error {
	//TODO implement me
	panic("implement me")
}

func (c checkData) UpdateCheck(ctx context.Context, check domain.Check) error {
	//TODO implement me
	panic("implement me")
}

func (c checkData) GetCheckById(ctx context.Context, checkId int64) (domain.Check, error) {
	//TODO implement me
	panic("implement me")
}

func (c checkData) ListChecks(ctx context.Context, pagination domain.Pagination, status *string) ([]domain.Check, error) {
	//TODO implement me
	panic("implement me")
}

func (c checkData) SubmitCheck(ctx context.Context, checkId int64, approved bool, comments string) error {
	//TODO implement me
	panic("implement me")
}

func (c checkData) BatchDeleteChecks(ctx context.Context, checkIds []int64) error {
	//TODO implement me
	panic("implement me")
}

func (c checkData) BatchSubmitChecks(ctx context.Context, checks []domain.Check) error {
	//TODO implement me
	panic("implement me")
}
