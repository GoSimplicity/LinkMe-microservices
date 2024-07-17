package data

import (
	"context"
	"errors"
	"github.com/GoSimplicity/LinkMe/app/linkme-check/domain"
	"github.com/GoSimplicity/LinkMe/app/linkme-check/internal/biz"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

type Check struct {
	ID        int64  `gorm:"primaryKey;autoIncrement"`
	PostID    int64  `gorm:"not null"`
	Content   string `gorm:"type:text;not null"`
	Title     string `gorm:"size:255;not null"`
	UserId    int64  `gorm:"column:user_id;index"`
	Status    string `gorm:"size:20;not null;default:'Pending'"`
	Remark    string `gorm:"type:text"`
	CreatedAt int64  `gorm:"column:created_at;type:bigint;not null"`
	UpdatedAt int64  `gorm:"column:updated_at;type:bigint;not null;index"`
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

func (c *checkData) getCurrentTime() int64 {
	return time.Now().UnixMilli()
}

func (c *checkData) CreateCheck(ctx context.Context, check domain.Check) (int64, error) {
	check.CreatedAt = c.getCurrentTime()
	check.UpdatedAt = c.getCurrentTime()
	dc := toDataCheck(check)
	if err := c.db.WithContext(ctx).Create(&dc).Error; err != nil {
		c.l.Error("failed to create check", zap.Error(err))
		return 0, err
	}
	return dc.ID, nil
}

func (c *checkData) DeleteCheck(ctx context.Context, checkId int64) error {
	now := c.getCurrentTime()
	if err := c.db.WithContext(ctx).Model(&Check{}).Where("id = ?", checkId).Update("deleted_at", now).Error; err != nil {
		c.l.Error("failed to delete check", zap.Error(err))
		return err
	}
	return nil
}

func (c *checkData) UpdateCheck(ctx context.Context, check domain.Check) error {
	check.UpdatedAt = c.getCurrentTime()
	dc := toDataCheck(check)
	if err := c.db.WithContext(ctx).Model(&Check{}).Where("id = ?", check.ID).Updates(dc).Error; err != nil {
		c.l.Error("failed to update check", zap.Error(err))
		return err
	}
	return nil
}

func (c *checkData) GetCheckById(ctx context.Context, checkId int64) (domain.Check, error) {
	var check Check
	if err := c.db.WithContext(ctx).First(&check, checkId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.Check{}, nil
		}
		c.l.Error("failed to get check by id", zap.Error(err))
		return domain.Check{}, err
	}
	return toDomainCheck(check), nil
}

func (c *checkData) ListChecks(ctx context.Context, pagination domain.Pagination, status *string) ([]domain.Check, error) {
	var checks []Check
	query := c.db.WithContext(ctx).Model(&Check{})

	if status != nil {
		query = query.Where("status = ?", *status)
	}
	intSize := int(*pagination.Size)
	intOffset := int(*pagination.Offset)
	if err := query.Limit(intSize).Offset(intOffset).Find(&checks).Error; err != nil {
		c.l.Error("failed to list checks", zap.Error(err))
		return nil, err
	}
	return toDomainSliceCheck(checks), nil
}

func (c *checkData) SubmitCheck(ctx context.Context, checkId int64, approved bool) error {
	status := "Rejected"
	if approved {
		status = "Approved"
	}
	if err := c.db.WithContext(ctx).Model(&Check{}).Where("id = ?", checkId).Update("status", status).Error; err != nil {
		c.l.Error("failed to submit check", zap.Error(err))
		return err
	}
	return nil
}

func (c *checkData) BatchDeleteChecks(ctx context.Context, checkIds []int64) error {
	now := c.getCurrentTime()
	if err := c.db.WithContext(ctx).Model(&Check{}).Where("id IN ?", checkIds).Update("deleted_at", now).Error; err != nil {
		c.l.Error("failed to batch delete checks", zap.Error(err))
		return err
	}
	return nil
}

func (c *checkData) BatchSubmitChecks(ctx context.Context, checks []domain.Check) error {
	tx := c.db.WithContext(ctx).Begin()
	for _, check := range checks {
		status := "Rejected"
		if check.Status == "Approved" {
			status = "Approved"
		}
		if err := tx.Model(&Check{}).Where("id = ?", check.ID).Update("status", status).Error; err != nil {
			tx.Rollback()
			c.l.Error("failed to batch submit checks", zap.Error(err))
			return err
		}
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		c.l.Error("failed to commit transaction", zap.Error(err))
		return err
	}
	return nil
}

// 转换函数
func toDomainCheck(check Check) domain.Check {
	return domain.Check{
		ID:        check.ID,
		PostID:    check.PostID,
		Status:    check.Status,
		Content:   check.Content,
		Title:     check.Title,
		Remark:    check.Remark,
		UserId:    check.UserId,
		CreatedAt: check.CreatedAt,
		UpdatedAt: check.UpdatedAt,
	}
}

func toDataCheck(check domain.Check) Check {
	return Check{
		ID:        check.ID,
		PostID:    check.PostID,
		Status:    check.Status,
		Content:   check.Content,
		Title:     check.Title,
		Remark:    check.Remark,
		UserId:    check.UserId,
		CreatedAt: check.CreatedAt,
		UpdatedAt: check.UpdatedAt,
	}
}

func toDomainSliceCheck(checks []Check) []domain.Check {
	domainChecks := make([]domain.Check, len(checks))
	for i, check := range checks {
		domainChecks[i] = toDomainCheck(check)
	}
	return domainChecks
}

func toDataSliceCheck(checks []domain.Check) []Check {
	dataChecks := make([]Check, len(checks))
	for i, check := range checks {
		dataChecks[i] = toDataCheck(check)
	}
	return dataChecks
}
