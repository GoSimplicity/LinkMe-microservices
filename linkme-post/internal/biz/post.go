package biz

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
	"linkme-post/api/post/v1"
	"linkme-post/domain"
	"linkme-post/internal/data"
)

type PostUsecase struct {
	repo data.PostRepo
	l    *zap.Logger
}

func NewPostUsecase(repo data.PostRepo, l *zap.Logger) *PostUsecase {
	return &PostUsecase{
		repo: repo,
		l:    l,
	}
}

func (uc *PostUsecase) CreatePost(ctx context.Context, dp domain.Post) (int64, error) {
	postId, err := uc.repo.Insert(ctx, fromDomainPost(dp))
	if err != nil {
		return -1, err
	}
	return postId, err
}

func (uc *PostUsecase) UpdatePost(ctx context.Context, req *post.UpdatePostRequest) (*post.UpdatePostReply, error) {
	// Your logic here
	return &post.UpdatePostReply{
		Code: 0,
		Msg:  "success",
	}, nil
}

func (uc *PostUsecase) DeletePost(ctx context.Context, req *post.DeletePostRequest) (*post.DeletePostReply, error) {
	// Your logic here
	return &post.DeletePostReply{
		Code: 0,
		Msg:  "success",
	}, nil
}

func (uc *PostUsecase) PublishPost(ctx context.Context, req *post.PublishPostRequest) (*post.PublishPostReply, error) {
	// Your logic here
	return &post.PublishPostReply{
		Code: 0,
		Msg:  "success",
	}, nil
}

func (uc *PostUsecase) WithdrawPost(ctx context.Context, req *post.WithdrawPostRequest) (*post.WithdrawPostReply, error) {
	// Your logic here
	return &post.WithdrawPostReply{
		Code: 0,
		Msg:  "success",
	}, nil
}

func (uc *PostUsecase) ListPost(ctx context.Context, req *post.ListPostRequest) (*post.ListPostReply, error) {
	// Your logic here
	return &post.ListPostReply{
		Code: 0,
		Msg:  "success",
		Data: nil, // fill this with actual data
	}, nil
}

func (uc *PostUsecase) ListPubPost(ctx context.Context, req *post.ListPubPostRequest) (*post.ListPubPostReply, error) {
	// Your logic here
	return &post.ListPubPostReply{
		Code: 0,
		Msg:  "success",
		Data: nil, // fill this with actual data
	}, nil
}

func (uc *PostUsecase) ListAdminPost(ctx context.Context, req *post.ListAdminPostRequest) (*post.ListAdminPostReply, error) {
	// Your logic here
	return &post.ListAdminPostReply{
		Code: 0,
		Msg:  "success",
		Data: nil, // fill this with actual data
	}, nil
}

func (uc *PostUsecase) DetailPost(ctx context.Context, req *post.DetailPostRequest) (*post.DetailPostReply, error) {
	// Your logic here
	return &post.DetailPostReply{
		Code: 0,
		Msg:  "success",
		Data: nil, // fill this with actual data
	}, nil
}

func (uc *PostUsecase) DetailPubPost(ctx context.Context, req *post.DetailPubPostRequest) (*post.DetailPubPostReply, error) {
	// Your logic here
	return &post.DetailPubPostReply{
		Code: 0,
		Msg:  "success",
		Data: nil, // fill this with actual data
	}, nil
}

func (uc *PostUsecase) DetailAdminPost(ctx context.Context, req *post.DetailAdminPostRequest) (*post.DetailAdminPostReply, error) {
	// Your logic here
	return &post.DetailAdminPostReply{
		Code: 0,
		Msg:  "success",
		Data: nil, // fill this with actual data
	}, nil
}

func (uc *PostUsecase) GetPostStats(ctx context.Context, _ *emptypb.Empty) (*post.GetPostStatsReply, error) {
	// Your logic here
	return &post.GetPostStatsReply{
		Code:  0,
		Msg:   "success",
		Count: 0, // fill this with actual data
	}, nil
}

func (uc *PostUsecase) LikePost(ctx context.Context, req *post.LikePostRequest) (*post.LikePostReply, error) {
	// Your logic here
	return &post.LikePostReply{
		Code: 0,
		Msg:  "success",
	}, nil
}

func (uc *PostUsecase) CollectPost(ctx context.Context, req *post.CollectPostRequest) (*post.CollectPostReply, error) {
	// Your logic here
	return &post.CollectPostReply{
		Code: 0,
		Msg:  "success",
	}, nil
}

// 将领域层对象转为dao层对象
func fromDomainPost(p domain.Post) data.Post {
	return data.Post{
		ID:         p.ID,
		Title:      p.Title,
		Content:    p.Content,
		CreatedAt:  p.CreateAt,
		UpdatedAt:  p.UpdatedAt,
		UserID:     p.UserID,
		Status:     p.Status,
		PlateID:    p.Plate.ID,
		LikeNum:    p.LikeNum,
		CollectNum: p.CollectNum,
		ViewNum:    p.ViewNum,
		Deleted:    p.Deleted,
		DeletedAt:  p.DeletedAt,
	}
}

// 将dao层对象转为领域层对象
func fromDomainSlicePost(post []data.Post) []domain.Post {
	domainPosts := make([]domain.Post, len(post)) // 创建与输入切片等长的domain.Post切片
	for i, repoPost := range post {
		domainPosts[i] = domain.Post{
			ID:         repoPost.ID,
			Title:      repoPost.Title,
			Content:    repoPost.Content,
			CreateAt:   repoPost.CreatedAt,
			UpdatedAt:  repoPost.UpdatedAt,
			Status:     repoPost.Status,
			Plate:      domain.Plate{ID: repoPost.PlateID},
			LikeNum:    repoPost.LikeNum,
			CollectNum: repoPost.CollectNum,
			ViewNum:    repoPost.ViewNum,
			Deleted:    repoPost.Deleted,
			DeletedAt:  repoPost.DeletedAt,
			UserID:     repoPost.UserID,
		}
	}
	return domainPosts
}

// 将dao层转化为领域层
func toDomainPost(post data.Post) domain.Post {
	return domain.Post{
		ID:         post.ID,
		Title:      post.Title,
		Content:    post.Content,
		CreateAt:   post.CreatedAt,
		UpdatedAt:  post.UpdatedAt,
		Status:     post.Status,
		Plate:      domain.Plate{ID: post.PlateID},
		UserID:     post.UserID,
		LikeNum:    post.LikeNum,
		CollectNum: post.CollectNum,
		ViewNum:    post.ViewNum,
		Deleted:    post.Deleted,
		DeletedAt:  post.DeletedAt,
	}
}
