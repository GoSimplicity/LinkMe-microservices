package data

import (
	"context"
	"errors"
	"fmt"
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-post/internal/biz"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var (
	ErrPostNotFound  = errors.New("post not found")
	ErrInvalidParams = errors.New("invalid parameters")
)

type postData struct {
	data *Data
	l    *zap.Logger
}

func NewPostData(data *Data, l *zap.Logger) biz.PostData {
	return &postData{
		data: data,
		l:    l,
	}
}

// Insert 创建一个新的帖子记录
func (p *postData) Insert(ctx context.Context, post biz.Post) (int64, error) {
	err := p.data.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 检查 plate_id 是否存在
		var count int64

		if err := tx.Model(&biz.Plate{}).Where("id = ?", post.PlateID).Count(&count).Error; err != nil {
			p.l.Error("failed to check plate existence", zap.Error(err))
			return err
		}

		if count == 0 {

			return errors.New("plate not found")
		}
		// 创建帖子
		if err := tx.Create(&post).Error; err != nil {
			p.l.Error("failed to insert post", zap.Error(err))
			return err
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return post.ID, nil
}

// UpdateById 通过Id更新帖子
func (p *postData) UpdateById(ctx context.Context, post biz.Post) error {
	if post.ID == 0 || post.UserID == 0 {
		return ErrInvalidParams
	}

	res := p.data.db.WithContext(ctx).Model(&biz.Post{}).Where("id = ? AND user_id = ?", post.ID, post.UserID).Updates(map[string]interface{}{
		"title":    post.Title,
		"content":  post.Content,
		"plate_id": post.PlateID,
		"status":   post.Status,
	})

	if res.Error != nil {
		p.l.Error("failed to update post", zap.Error(res.Error))
		return res.Error
	}

	if res.RowsAffected == 0 {
		return ErrPostNotFound
	}

	return nil
}

// UpdateStatus 更新帖子状态
func (p *postData) UpdateStatus(ctx context.Context, post biz.Post) error {
	res := p.data.db.WithContext(ctx).Model(&biz.Post{}).Where("id = ?", post.ID).Updates(map[string]interface{}{
		"status": post.Status,
	})

	if res.Error != nil {
		p.l.Error("failed to update post status", zap.Error(res.Error))
		return res.Error
	}

	if res.RowsAffected == 0 {
		return ErrPostNotFound
	}

	return nil
}

// GetById 根据ID获取一个帖子记录
func (p *postData) GetById(ctx context.Context, postId int64, uid int64) (biz.Post, error) {
	var post biz.Post

	err := p.data.db.WithContext(ctx).Where("user_id = ? AND id = ?", uid, postId).First(&post).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			p.l.Debug("post not found", zap.Int64("id", postId), zap.Int64("user_id", uid))
			return biz.Post{}, ErrPostNotFound
		}

		p.l.Error("failed to get post", zap.Error(err), zap.Int64("id", postId), zap.Int64("user_id", uid))

		return biz.Post{}, err
	}

	return post, nil
}

// GetPubById 根据ID获取一个已发布的帖子记录
func (p *postData) GetPubById(ctx context.Context, postId int64) (biz.ListPubPost, error) {
	var post biz.ListPubPost

	// 设置查询超时时间
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	status := biz.Published
	// 设置查询过滤器，只查找状态为已发布的帖子
	filter := bson.M{
		"id":     postId,
		"status": status,
	}

	// 在MongoDB的posts集合中查找记录
	err := p.data.mongo.Database("linkme").Collection("posts").FindOne(ctx, filter).Decode(&post)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			p.l.Debug("published post not found", zap.Error(err))
			return biz.ListPubPost{}, ErrPostNotFound
		}
		p.l.Error("failed to get published post", zap.Error(err))
		return biz.ListPubPost{}, err
	}

	return post, nil
}

// ListPub 查询公开帖子列表
func (p *postData) ListPub(ctx context.Context, pagination biz.Pagination) ([]biz.ListPubPost, error) {
	status := biz.Published
	var posts []biz.ListPubPost

	// 设置查询超时时间
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// 指定数据库与集合
	collection := p.data.mongo.Database("linkme").Collection("posts")

	filter := bson.M{
		"status": status,
	}
	// 设置分页查询参数
	opts := options.FindOptions{
		Skip:  pagination.Offset,
		Limit: pagination.Size,
	}

	cursor, err := collection.Find(ctx, filter, &opts)
	if err != nil {
		p.l.Error("database query failed", zap.Error(err))
		return nil, err
	}

	// 将获取到的查询结果解码到posts结构体中
	if err = cursor.All(ctx, &posts); err != nil {
		p.l.Error("failed to decode query results", zap.Error(err))
		return nil, err
	}

	if len(posts) == 0 {
		p.l.Debug("query returned no results")
	}

	return posts, nil
}

// List 查询作者帖子列表
func (p *postData) List(ctx context.Context, pagination biz.Pagination) ([]biz.Post, error) {
	var posts []biz.Post

	if err := p.data.db.WithContext(ctx).Where("user_id = ?", pagination.Uid).
		Limit(int(*pagination.Size)).Offset(int(*pagination.Offset)).Find(&posts).Error; err != nil {

		p.l.Error("find post list failed", zap.Error(err))

		return nil, err
	}

	return posts, nil
}

// DeleteById 通过id删除帖子
func (p *postData) DeleteById(ctx context.Context, post biz.Post) error {
	// 使用事务来确保操作的原子性
	tx := p.data.db.WithContext(ctx).Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 更新帖子的状态为已删除，并使用 GORM 的软删除功能
	if err := tx.Model(&biz.Post{}).Where("id = ?", post.ID).Updates(map[string]interface{}{
		"status": biz.Deleted,
	}).Delete(&biz.Post{}).Error; err != nil {
		tx.Rollback()

		p.l.Error("failed to delete post", zap.Error(err))

		return err
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()

		p.l.Error("failed to commit transaction", zap.Error(err))

		return err
	}

	return nil
}

func (p *postData) CreatePlate(ctx context.Context, plate biz.Plate) error {
	if err := p.data.db.WithContext(ctx).Create(&plate).Error; err != nil {
		p.l.Error("failed to insert plate", zap.Error(err))
		return fmt.Errorf("failed to create plate: %w", err)
	}

	return nil
}

func (p *postData) ListPlates(ctx context.Context, pagination biz.Pagination) ([]biz.Plate, error) {
	var plates []biz.Plate

	// 使用 offset 和 limit 分页，避免 Limit 方法被误用
	if err := p.data.db.WithContext(ctx).
		Offset(int(*pagination.Offset)).
		Limit(int(*pagination.Size)).
		Find(&plates).Error; err != nil {
		p.l.Error("failed to list plates", zap.Error(err))
		return nil, fmt.Errorf("failed to list plates: %w", err)
	}

	return plates, nil
}

func (p *postData) UpdatePlate(ctx context.Context, plate biz.Plate) error {
	updates := map[string]interface{}{
		"name": plate.Name,
	}

	if err := p.data.db.WithContext(ctx).
		Model(&biz.Plate{}).
		Where("id = ?", plate.ID).
		Updates(updates).Error; err != nil {
		p.l.Error("failed to update plate", zap.Error(err))
		return fmt.Errorf("failed to update plate: %w", err)
	}

	return nil
}

func (p *postData) DeletePlate(ctx context.Context, plateId int64) error {
	if err := p.data.db.WithContext(ctx).
		Where("id = ?", plateId).
		Delete(&biz.Plate{}).Error; err != nil {
		p.l.Error("failed to delete plate", zap.Error(err))
		return fmt.Errorf("failed to delete plate: %w", err)
	}

	return nil
}
