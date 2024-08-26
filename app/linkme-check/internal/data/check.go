package data

import (
	"context"
	"errors"
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-check/internal/biz"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

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

func (c *checkData) Create(ctx context.Context, check biz.Check) (int64, error) {
	// 参数验证
	if check.PostID == 0 || check.Content == "" || check.Title == "" || check.UserID == 0 {
		return 0, errors.New("invalid input: missing required fields")
	}

	// 使用事务处理以确保数据一致性
	err := c.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&check).Error; err != nil {
			c.l.Error("failed to create check", zap.Error(err))
			return err
		}
		return nil
	})

	if err != nil {
		return 0, err
	}

	return check.ID, nil
}

func (c *checkData) UpdateStatus(ctx context.Context, check biz.Check) error {
	if check.ID == 0 {
		return errors.New("invalid input: missing required fields")
	}

	// 使用事务处理更新操作
	err := c.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		result := tx.Model(&biz.Check{}).Where("id = ?", check.ID).Update("status", check.Status)
		if result.Error != nil {
			c.l.Error("failed to update check status", zap.Error(result.Error))
			return result.Error
		}

		if result.RowsAffected == 0 {
			return errors.New("no records updated")
		}
		return nil
	})

	return err
}

func (c *checkData) Delete(ctx context.Context, checkId int64, uid int64) error {
	// 验证 checkId 是否有效
	if checkId == 0 {
		return errors.New("invalid input: missing check ID")
	}

	err := c.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 执行软删除操作
		result := tx.Where("id = ? AND user_id = ?", checkId, uid).Delete(&biz.Check{})
		if result.Error != nil {
			c.l.Error("failed to soft delete check", zap.Error(result.Error))
			return result.Error
		}

		// 检查是否有记录被删除
		if result.RowsAffected == 0 {
			return errors.New("no records deleted")
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (c *checkData) ListChecks(ctx context.Context, pagination biz.Pagination) ([]biz.Check, error) {
	var checks []biz.Check

	// 在查询时，避免将 Size 或 Offset 的指针传递为空值导致的 panic
	size := int(*pagination.Size)
	offset := int(*pagination.Offset)

	err := c.db.WithContext(ctx).
		Limit(size).
		Offset(offset).
		Find(&checks).Error

	if err != nil {
		c.l.Error("failed to find all checks", zap.Error(err))
		return nil, err
	}

	return checks, nil
}

func (c *checkData) FindByID(ctx context.Context, checkId int64) (biz.Check, error) {
	var check biz.Check

	err := c.db.WithContext(ctx).Where("id = ?", checkId).First(&check).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return biz.Check{}, nil
		}
		c.l.Error("failed to find check by ID", zap.Error(err))
		return biz.Check{}, err
	}

	return check, nil
}

func (c *checkData) FindByPostId(ctx context.Context, postId uint) (biz.Check, error) {
	var check biz.Check

	err := c.db.WithContext(ctx).Where("post_id = ?", postId).First(&check).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return biz.Check{}, nil
		}
		c.l.Error("failed to find check by post ID", zap.Error(err))
		return biz.Check{}, err
	}

	return check, nil
}

func (c *checkData) GetCheckCount(ctx context.Context) (int64, error) {
	var count int64

	err := c.db.WithContext(ctx).Model(&biz.Check{}).Where("status = ?", biz.UnderReview).Count(&count).Error
	if err != nil {
		c.l.Error("failed to get check count", zap.Error(err))
		return -1, err
	}

	return count, nil
}
