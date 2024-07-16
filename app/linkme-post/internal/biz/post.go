package biz

import (
	"context"
	"errors"
	"github.com/GoSimplicity/LinkMe-monorepo/api/post/v1"
	"github.com/GoSimplicity/LinkMe/app/linkme-post/domain"
)

var (
	ErrPostNotFound  = errors.New("post not found")
	ErrInvalidParams = errors.New("invalid parameters")
	ErrSyncFailed    = errors.New("sync failed")
)

type PostData interface {
	CreatePost(ctx context.Context, dp domain.Post) (int64, error)                         // 创建一个新的帖子记录
	CreatePubPost(ctx context.Context, dp domain.Post) (int64, error)                      // 创建一个新的公开帖子记录
	DeletePost(ctx context.Context, postId int64, uid int64) error                         // 删除一个帖子
	UpdatePost(ctx context.Context, dp domain.Post) error                                  // 根据ID更新一个帖子记录
	UpdatePostStatus(ctx context.Context, dp domain.Post) error                            // 更新帖子的状态
	GetPost(ctx context.Context, postId int64, uid int64) (domain.Post, error)             // 根据ID获取一个帖子记录
	GetPubPost(ctx context.Context, id int64) (domain.Post, error)                         // 根据ID获取一个已发布的帖子记录
	ListPosts(ctx context.Context, pagination domain.Pagination) ([]domain.Post, error)    // 获取个人的帖子记录列表
	ListPubPosts(ctx context.Context, pagination domain.Pagination) ([]domain.Post, error) // 获取已发布的帖子记录列表
	SyncPost(ctx context.Context, dp domain.Post) (int64, error)                           // 用于同步帖子记录
}

type PostBiz struct {
	postData PostData
}

func NewPostBiz(postData PostData) *PostBiz {
	return &PostBiz{
		postData: postData,
	}
}

func (pb *PostBiz) CreatePost(ctx context.Context, dp domain.Post) (int64, error) {
	dp.Status = domain.Draft
	postId, err := pb.postData.CreatePost(ctx, dp)
	if err != nil {
		return -1, err
	}
	return postId, err
}

func (pb *PostBiz) UpdatePost(ctx context.Context, dp domain.Post) error {
	dp.Status = domain.Draft
	err := pb.postData.UpdatePost(ctx, dp)
	if err != nil {
		return err
	}
	return err
}

func (pb *PostBiz) UpdatePostStatus(ctx context.Context, dp domain.Post) error {
	return pb.postData.UpdatePostStatus(ctx, dp)
}

func (pb *PostBiz) GetPost(ctx context.Context, postId int64, uid int64) (domain.Post, error) {
	return pb.postData.GetPost(ctx, postId, uid)
}

func (pb *PostBiz) GetPubPost(ctx context.Context, postId int64) (domain.Post, error) {
	return pb.postData.GetPubPost(ctx, postId)
}

func (pb *PostBiz) ListPost(ctx context.Context, pagination domain.Pagination) ([]domain.Post, error) {
	offset := int64(pagination.Page-1) * *pagination.Size
	pagination.Offset = &offset
	return pb.postData.ListPosts(ctx, pagination)
}

func (pb *PostBiz) ListPubPost(ctx context.Context, pagination domain.Pagination) ([]domain.Post, error) {
	offset := int64(pagination.Page-1) * *pagination.Size
	pagination.Offset = &offset
	return pb.postData.ListPubPosts(ctx, pagination)
}

func (pb *PostBiz) DeletePost(ctx context.Context, postId int64, uid int64) error {
	return pb.postData.DeletePost(ctx, postId, uid)
}

func (pb *PostBiz) CreatePubPost(ctx context.Context, dp domain.Post) (int64, error) {
	// TODO 暂时保留
	return 0, nil
}

func (pb *PostBiz) LikePost(ctx context.Context, req *post.LikePostRequest) (*post.LikePostReply, error) {
	// TODO 暂时保留
	return nil, nil
}

func (pb *PostBiz) CollectPost(ctx context.Context, req *post.CollectPostRequest) (*post.CollectPostReply, error) {
	// TODO 暂时保留
	return nil, nil
}
