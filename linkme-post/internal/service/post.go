package service

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	pb "linkme-post/api/post/v1"
	"linkme-post/domain"
	"linkme-post/internal/biz"
)

type PostService struct {
	pb.UnimplementedPostServer
	biz *biz.PostUsecase
}

func NewPostService(biz *biz.PostUsecase) *PostService {
	return &PostService{
		biz: biz,
	}
}

func (s *PostService) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.CreatePostReply, error) {
	post, err := s.biz.CreatePost(ctx, domain.Post{
		Title:   req.Title,
		Content: req.Content,
		Plate:   domain.Plate{ID: req.PlateId},
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreatePostReply{
		Code: 0,
		Msg:  "create post success",
		Data: string(post),
	}, nil
}
func (s *PostService) UpdatePost(ctx context.Context, req *pb.UpdatePostRequest) (*pb.UpdatePostReply, error) {
	return &pb.UpdatePostReply{}, nil
}
func (s *PostService) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*pb.DeletePostReply, error) {
	return &pb.DeletePostReply{}, nil
}
func (s *PostService) PublishPost(ctx context.Context, req *pb.PublishPostRequest) (*pb.PublishPostReply, error) {
	return &pb.PublishPostReply{}, nil
}
func (s *PostService) WithdrawPost(ctx context.Context, req *pb.WithdrawPostRequest) (*pb.WithdrawPostReply, error) {
	return &pb.WithdrawPostReply{}, nil
}
func (s *PostService) ListPost(ctx context.Context, req *pb.ListPostRequest) (*pb.ListPostReply, error) {
	return &pb.ListPostReply{}, nil
}
func (s *PostService) ListPubPost(ctx context.Context, req *pb.ListPubPostRequest) (*pb.ListPubPostReply, error) {
	return &pb.ListPubPostReply{}, nil
}
func (s *PostService) ListAdminPost(ctx context.Context, req *pb.ListAdminPostRequest) (*pb.ListAdminPostReply, error) {
	return &pb.ListAdminPostReply{}, nil
}
func (s *PostService) DetailPost(ctx context.Context, req *pb.DetailPostRequest) (*pb.DetailPostReply, error) {
	return &pb.DetailPostReply{}, nil
}
func (s *PostService) DetailPubPost(ctx context.Context, req *pb.DetailPubPostRequest) (*pb.DetailPubPostReply, error) {
	return &pb.DetailPubPostReply{}, nil
}
func (s *PostService) DetailAdminPost(ctx context.Context, req *pb.DetailAdminPostRequest) (*pb.DetailAdminPostReply, error) {
	return &pb.DetailAdminPostReply{}, nil
}
func (s *PostService) GetPostStats(ctx context.Context, req *empty.Empty) (*pb.GetPostStatsReply, error) {
	return &pb.GetPostStatsReply{}, nil
}
func (s *PostService) LikePost(ctx context.Context, req *pb.LikePostRequest) (*pb.LikePostReply, error) {
	return &pb.LikePostReply{}, nil
}
func (s *PostService) CollectPost(ctx context.Context, req *pb.CollectPostRequest) (*pb.CollectPostReply, error) {
	return &pb.CollectPostReply{}, nil
}
