package biz

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type CheckData interface {
	Create(ctx context.Context, check Check) error                          // 创建审核记录
	UpdateStatus(ctx context.Context, check Check) error                    // 更新审核状态
	Delete(ctx context.Context, checkId int64, uid int64) error             // 更新审核状态
	ListChecks(ctx context.Context, pagination Pagination) ([]Check, error) // 获取审核列表
	FindByID(ctx context.Context, checkId int64) (Check, error)             // 通过ID获取审核记录
	FindByPostId(ctx context.Context, postId uint) (Check, error)           // 通过帖子ID获取审核记录
	GetCheckCount(ctx context.Context) (int64, error)                       // 获取审核记录数量
}

const (
	UnderReview uint8 = iota
	Approved
	UnApproved
)

type Check struct {
	ID        int64        `gorm:"primaryKey;autoIncrement"`         // 审核ID
	PostID    int64        `gorm:"not null"`                         // 帖子ID
	Content   string       `gorm:"type:text;not null"`               // 审核内容
	Title     string       `gorm:"size:255;not null"`                // 审核标签
	UserID    int64        `gorm:"column:user_id;index"`             // 提交审核的用户ID
	Status    uint8        `gorm:"default:0"`                        // 审核状态
	Remark    string       `gorm:"type:text"`                        // 审核备注
	CreatedAt time.Time    `gorm:"column:created_at;autoCreateTime"` // 创建时间
	UpdatedAt time.Time    `gorm:"column:updated_at;autoUpdateTime"` // 更新时间
	DeletedAt sql.NullTime `gorm:"index"`                            // 删除时间
}

type Pagination struct {
	Page int    // 当前页码
	Size *int64 // 每页数据
	Uid  int64
	// 以下字段通常在服务端内部使用，不需要客户端传递
	Offset *int64 // 数据偏移量
	Total  *int64 // 总数据量
}

type CheckBiz struct {
	CheckData CheckData
}

func NewCheckBiz(CheckData CheckData) *CheckBiz {
	return &CheckBiz{CheckData: CheckData}
}

func (c *CheckBiz) CreateCheck(ctx context.Context, check Check) error {
	check.Status = UnderReview
	return c.CheckData.Create(ctx, check)
}

func (c *CheckBiz) UpdateStatus(ctx context.Context, checkID int64, remark string, status uint8) error {
	check, err := c.CheckData.FindByID(ctx, checkID)
	if err != nil {
		return fmt.Errorf("获取审核详情失败: %w", err)
	}

	if check.Status != UnderReview {
		return errors.New("请勿重复提交")
	}

	check.Remark = remark
	check.Status = status
	return c.CheckData.UpdateStatus(ctx, check)
}

func (c *CheckBiz) DeleteCheck(ctx context.Context, checkID int64, userId int64) error {
	return c.CheckData.Delete(ctx, checkID, userId)
}

func (c *CheckBiz) GetCheck(ctx context.Context, checkID int64) (Check, error) {
	return c.CheckData.FindByID(ctx, checkID)
}

func (c *CheckBiz) ListChecks(ctx context.Context, pagination Pagination) ([]Check, error) {
	// 计算偏移量
	offset := int64(pagination.Page-1) * *pagination.Size
	pagination.Offset = &offset
	return c.CheckData.ListChecks(ctx, pagination)
}
