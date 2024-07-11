package data

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"linkme-post/domain"
	"time"
)

var (
	ErrPostNotFound  = errors.New("post not found")
	ErrInvalidParams = errors.New("invalid parameters")
	ErrSyncFailed    = errors.New("sync failed")
)

type PostRepo interface {
	Insert(ctx context.Context, post Post) (int64, error)                      // 创建一个新的帖子记录
	UpdateById(ctx context.Context, post Post) error                           // 根据ID更新一个帖子记录
	Sync(ctx context.Context, post Post) (int64, error)                        // 用于同步帖子记录
	UpdateStatus(ctx context.Context, post Post) error                         // 更新帖子的状态
	GetByAuthor(ctx context.Context, postId int64, uid int64) (Post, error)    // 根据作者ID获取帖子记录
	GetById(ctx context.Context, id int64, uid int64) (Post, error)            // 根据ID获取一个帖子记录
	GetPubById(ctx context.Context, id int64) (Post, error)                    // 根据ID获取一个已发布的帖子记录
	ListPub(ctx context.Context, pagination domain.Pagination) ([]Post, error) // 获取已发布的帖子记录列表
	List(ctx context.Context, pagination domain.Pagination) ([]Post, error)    // 获取个人的帖子记录列表
	DeleteById(ctx context.Context, post domain.Post) error
	ListAllPost(ctx context.Context, pagination domain.Pagination) ([]Post, error)
	GetPost(ctx context.Context, id int64) (Post, error)
	GetPostCount(ctx context.Context) (int64, error)
}

type postRepoImpl struct {
	client *mongo.Client
	l      *zap.Logger
	db     *gorm.DB
}

type Post struct {
	ID         int64  `gorm:"primaryKey;autoIncrement"`
	Title      string `gorm:"size:255;not null"`                         // 文章标题
	Content    string `gorm:"type:text;not null"`                        // 文章内容
	CreatedAt  int64  `gorm:"column:created_at;autoCreateTime;not null"` // 创建时间
	UpdatedAt  int64  `gorm:"column:updated_at;autoUpdateTime;not null"` // 更新时间
	DeletedAt  int64  `gorm:"column:deleted_at;index"`                   // 删除时间
	Deleted    bool   `gorm:"column:deleted;default:false"`              // 是否删除
	Status     string `gorm:"size:20;default:'draft'"`                   // 文章状态，如草稿、发布等
	UserID     int64  `gorm:"column:user_id;index"`                      // 用户ID
	PlateID    int64  `gorm:"index"`                                     // 关联板块表的外键
	LikeNum    int64  `gorm:"column:like_num;default:0"`                 // 点赞数
	CollectNum int64  `gorm:"column:collect_num;default:0"`              // 收藏数
	ViewNum    int64  `gorm:"column:view_num;default:0"`                 // 浏览数
	Plate      Plate  `gorm:"foreignKey:PlateID"`                        // 板块关系
}

type Plate struct {
	ID          int64  `gorm:"primaryKey;autoIncrement"`         // 板块ID
	Name        string `gorm:"size:255;not null;uniqueIndex"`    // 板块名称
	Description string `gorm:"type:text"`                        // 板块描述
	CreatedAt   int64  `gorm:"column:created_at;autoCreateTime"` // 创建时间
	UpdatedAt   int64  `gorm:"column:updated_at;autoUpdateTime"` // 更新时间
	DeletedAt   int64  `gorm:"column:deleted_at;index"`          // 删除时间
	Deleted     bool   `gorm:"column:deleted;default:false"`     // 是否删除
	UserID      int64  `gorm:"column:uid;index"`                 // 板主ID
	Posts       []Post `gorm:"foreignKey:PlateID"`               // 帖子关系
}

func NewPostRepo(db *gorm.DB, l *zap.Logger, client *mongo.Client) PostRepo {
	return &postRepoImpl{
		client: client,
		l:      l,
		db:     db,
	}
}

// 获取当前时间的时间戳
func (p *postRepoImpl) getCurrentTime() int64 {
	return time.Now().UnixMilli()
}

// Insert 创建一个新的帖子记录
func (p *postRepoImpl) Insert(ctx context.Context, post Post) (int64, error) {
	now := p.getCurrentTime()
	post.CreatedAt = now
	post.UpdatedAt = now
	// 检查 plate_id 是否存在
	var count int64
	if err := p.db.WithContext(ctx).Model(&Plate{}).Where("id = ?", post.PlateID).Count(&count).Error; err != nil {
		p.l.Error("failed to check plate existence", zap.Error(err))
		return -1, err
	}
	if count == 0 {
		return -1, errors.New("plate not found")
	}
	// 创建帖子
	if err := p.db.WithContext(ctx).Create(&post).Error; err != nil {
		p.l.Error("failed to insert post", zap.Error(err))
		return -1, err
	}
	return post.ID, nil
}

// UpdateById 通过Id更新帖子
func (p *postRepoImpl) UpdateById(ctx context.Context, post Post) error {
	if post.ID == 0 || post.UserID == 0 {
		return ErrInvalidParams
	}
	var existingPost Post
	if err := p.db.WithContext(ctx).First(&existingPost, "id = ? AND author_id = ?", post.ID, post.UserID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrPostNotFound
		}
		p.l.Error("failed to find post", zap.Error(err))
		return err
	}
	// 检查是否有任何变化
	if existingPost.Title == post.Title &&
		existingPost.Content == post.Content &&
		existingPost.Status == post.Status &&
		existingPost.PlateID == post.PlateID {
		return errors.New("no changes") // 没有变化，不执行更新
	}
	now := p.getCurrentTime()
	updatedPost := map[string]any{
		"title":      post.Title,
		"content":    post.Content,
		"status":     post.Status,
		"plate_id":   post.PlateID,
		"updated_at": now,
	}
	res := p.db.WithContext(ctx).Model(&Post{}).Where("id = ? AND author_id = ?", post.ID, post.UserID).Updates(updatedPost)
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
func (p *postRepoImpl) UpdateStatus(ctx context.Context, post Post) error {
	now := p.getCurrentTime()
	if err := p.db.WithContext(ctx).Model(&Post{}).Where("id = ?", post.ID).
		Updates(map[string]any{
			"status":     post.Status,
			"updated_at": now,
		}).Error; err != nil {
		p.l.Error("failed to update post status", zap.Error(err))
		return err
	}
	return nil
}

// Sync 同步线上库(MongoDB)与制作库(MySQL)
func (p *postRepoImpl) Sync(ctx context.Context, post Post) (int64, error) {
	// 获取当前时间
	now := p.getCurrentTime()
	post.UpdatedAt = now
	// 获取 MySQL 中的帖子
	var mysqlPost Post
	err := p.db.WithContext(ctx).Where("id = ?", post.ID).First(&mysqlPost).Error
	if err != nil {
		return -1, err
	}
	// 检查帖子是否已存在于 MongoDB
	exists, err := p.checkPostExistsInMongoDB(ctx, post.ID)
	if err != nil {
		p.l.Error("failed to check post existence in MongoDB", zap.Error(err))
		return -1, err
	}
	if post.Status == domain.Published {
		if exists {
			// MongoDB 中已存在相同 ID 的文章，不执行同步
			return -1, ErrSyncFailed
		}
		// 插入帖子到 MongoDB
		if er := p.insertPostToMongoDB(ctx, mysqlPost); err != nil {
			p.l.Error("failed to insert post to MongoDB", zap.Error(er))
			return -1, er
		}
	} else {
		if exists {
			// 删除 MongoDB 中的帖子
			if er := p.deletePostFromMongoDB(ctx, post.ID); er != nil {
				p.l.Error("failed to delete post from MongoDB", zap.Error(er))
				return -1, er
			}
		}
	}
	return mysqlPost.ID, nil
}

// GetById 根据ID获取一个帖子记录
func (p *postRepoImpl) GetById(ctx context.Context, id int64, uid int64) (Post, error) {
	var post Post
	err := p.db.WithContext(ctx).Where("author_id = ? AND id = ? AND deleted = ?", uid, id, false).First(&post).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		p.l.Debug("post not found", zap.Error(err))
		return Post{}, ErrPostNotFound
	}
	return post, err
}

// GetPubById 根据ID获取一个已发布的帖子记录
func (p *postRepoImpl) GetPubById(ctx context.Context, id int64) (Post, error) {
	var post Post
	// 设置查询超时时间
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	status := domain.Published
	// 设置查询过滤器，只查找状态为已发布的帖子
	filter := bson.M{
		"id":     id,
		"status": status,
	}
	// 在MongoDB的posts集合中查找记录
	err := p.client.Database("linkme").Collection("posts").FindOne(ctx, filter).Decode(&post)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			p.l.Debug("published post not found", zap.Error(err))
			return Post{}, ErrPostNotFound
		}
		p.l.Error("failed to get published post", zap.Error(err))
		return Post{}, err
	}
	return post, nil
}

// GetByAuthor 根据作者ID获取帖子记录
func (p *postRepoImpl) GetByAuthor(ctx context.Context, postId int64, uid int64) (Post, error) {
	var post Post
	err := p.db.WithContext(ctx).Where("id = ? AND author_id = ? AND deleted = ?", postId, uid, false).Find(&post).Error
	if err != nil {
		p.l.Error("failed to get posts by author", zap.Error(err))
		return Post{}, err
	}
	return post, nil
}

// List 查询作者帖子列表
func (p *postRepoImpl) List(ctx context.Context, pagination domain.Pagination) ([]Post, error) {
	var posts []Post
	intSize := int(*pagination.Size)
	intOffset := int(*pagination.Offset)
	if err := p.db.WithContext(ctx).Where("author_id = ? AND deleted = ?", pagination.Uid, false).Limit(intSize).Offset(intOffset).Find(&posts).Error; err != nil {
		p.l.Error("find post list failed", zap.Error(err))
		return nil, err
	}
	return posts, nil
}

// ListPub 查询公开帖子列表
func (p *postRepoImpl) ListPub(ctx context.Context, pagination domain.Pagination) ([]Post, error) {
	status := domain.Published
	// 设置查询超时时间
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	// 指定数据库与集合
	collection := p.client.Database("linkme").Collection("posts")
	filter := bson.M{
		"status": status,
	}
	// 设置分页查询参数
	opts := options.FindOptions{
		Skip:  pagination.Offset,
		Limit: pagination.Size,
	}
	var posts []Post
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

// DeleteById 通过id删除帖子
func (p *postRepoImpl) DeleteById(ctx context.Context, post domain.Post) error {
	now := p.getCurrentTime()
	// 使用事务来确保操作的原子性
	tx := p.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	// 更新帖子的删除时间
	if err := tx.Model(Post{}).Where("id = ?", post.ID).Update("deleted_at", now).Update("status", domain.Deleted).Update("deleted", true).Error; err != nil {
		tx.Rollback()
		p.l.Error("failed to update post deletion time", zap.Error(err))
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

func (p *postRepoImpl) ListAllPost(ctx context.Context, pagination domain.Pagination) ([]Post, error) {
	var posts []Post
	intSize := int(*pagination.Size)
	intOffset := int(*pagination.Offset)
	if err := p.db.WithContext(ctx).Where("deleted = ?", false).Limit(intSize).Offset(intOffset).Find(&posts).Error; err != nil {
		p.l.Error("find post list failed", zap.Error(err))
		return nil, err
	}
	return posts, nil
}

func (p *postRepoImpl) GetPost(ctx context.Context, id int64) (Post, error) {
	var post Post
	err := p.db.WithContext(ctx).Where("id = ? AND deleted = ?", id, false).Find(&post).Error
	if err != nil {
		p.l.Error("find post failed", zap.Error(err))
		return Post{}, err
	}
	return post, nil
}

func (p *postRepoImpl) GetPostCount(ctx context.Context) (int64, error) {
	var count int64
	if err := p.db.WithContext(ctx).Model(Post{}).Count(&count).Error; err != nil {
		p.l.Error("failed to get post count", zap.Error(err))
		return -1, err
	}
	return count, nil
}

// checkPostExistsInMongoDB 检查帖子是否已存在于 MongoDB
func (p *postRepoImpl) checkPostExistsInMongoDB(ctx context.Context, postID int64) (bool, error) {
	var post Post
	err := p.client.Database("linkme").Collection("posts").FindOne(ctx, bson.M{"id": postID}).Decode(&post)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// insertPostToMongoDB 将帖子插入到 MongoDB
func (p *postRepoImpl) insertPostToMongoDB(ctx context.Context, post Post) error {
	_, err := p.client.Database("linkme").Collection("posts").InsertOne(ctx, post)
	return err
}

// deletePostFromMongoDB 从 MongoDB 中删除帖子
func (p *postRepoImpl) deletePostFromMongoDB(ctx context.Context, postID int64) error {
	_, err := p.client.Database("linkme").Collection("posts").DeleteOne(ctx, bson.M{"id": postID})
	return err
}
