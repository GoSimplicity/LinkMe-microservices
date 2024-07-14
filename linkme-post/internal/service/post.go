package service

import (
	"context"
	userpb "github.com/GoSimplicity/LinkMe-monorepo/api/user/v1"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/golang/protobuf/ptypes/empty"
	clientv3 "go.etcd.io/etcd/client/v3"
	pb "linkme-post/api/post/v1"
	"linkme-post/domain"
	"linkme-post/internal/biz"
	"time"
)

type PostService struct {
	pb.UnimplementedPostServer
	userClient userpb.UserClient
	biz        *biz.PostBiz
}

func NewPostService(biz *biz.PostBiz) *PostService {
	return &PostService{
		biz: biz,
	}
}

func (s *PostService) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.CreatePostReply, error) {
	userId, err := s.getUserId(ctx)
	if err != nil {
		return &pb.CreatePostReply{
			Code: 1,
			Msg:  err.Error(),
		}, err
	}
	postId, err := s.biz.CreatePost(ctx, domain.Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  userId,
		Plate:   domain.Plate{ID: req.PlateId},
	})
	if err != nil {
		return &pb.CreatePostReply{
			Code: 1,
			Msg:  err.Error(),
		}, err
	}
	return &pb.CreatePostReply{
		Code: 0,
		Msg:  "create post success",
		Data: postId,
	}, nil
}
func (s *PostService) UpdatePost(ctx context.Context, req *pb.UpdatePostRequest) (*pb.UpdatePostReply, error) {
	userId, err := s.getUserId(ctx)
	if err != nil {
		return &pb.UpdatePostReply{
			Code: 1,
			Msg:  err.Error(),
		}, err
	}
	err = s.biz.UpdatePost(ctx, domain.Post{
		ID:      req.PostId,
		Title:   req.Title,
		Content: req.Content,
		UserID:  userId,
		Plate:   domain.Plate{ID: req.PlateId},
	})
	if err != nil {
		return &pb.UpdatePostReply{
			Code: 1,
			Msg:  err.Error(),
		}, err
	}
	return &pb.UpdatePostReply{
		Code: 0,
		Msg:  "update post success",
	}, nil
}
func (s *PostService) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*pb.DeletePostReply, error) {
	userId, err := s.getUserId(ctx)
	if err != nil {
		return &pb.DeletePostReply{
			Code: 1,
			Msg:  err.Error(),
		}, err
	}
	err = s.biz.DeletePost(ctx, req.PostId, userId)
	if err != nil {
		return &pb.DeletePostReply{
			Code: 1,
			Msg:  err.Error(),
		}, err
	}
	return &pb.DeletePostReply{
		Code: 0,
		Msg:  "delete post success",
	}, nil
}
func (s *PostService) PublishPost(ctx context.Context, req *pb.PublishPostRequest) (*pb.PublishPostReply, error) {
	// TODO 暂时保留
	return &pb.PublishPostReply{
		Code: 0,
		Msg:  "publish post success",
	}, nil
}
func (s *PostService) WithdrawPost(ctx context.Context, req *pb.WithdrawPostRequest) (*pb.WithdrawPostReply, error) {
	userId, err := s.getUserId(ctx)
	if err != nil {
		return &pb.WithdrawPostReply{
			Code: 1,
			Msg:  err.Error(),
		}, err
	}
	dp := domain.Post{
		ID:     req.PostId,
		UserID: userId,
		Status: domain.Withdrawn,
	}
	err = s.biz.UpdatePostStatus(ctx, dp)
	if err != nil {
		return &pb.WithdrawPostReply{
			Code: 1,
			Msg:  err.Error(),
		}, err
	}
	return &pb.WithdrawPostReply{
		Code: 0,
		Msg:  "withdraw post success",
	}, nil
}
func (s *PostService) ListPost(ctx context.Context, req *pb.ListPostRequest) (*pb.ListPostReply, error) {
	userId, err := s.getUserId(ctx)
	if err != nil {
		return &pb.ListPostReply{
			Code: 1,
			Msg:  err.Error(),
		}, err
	}
	pagination := domain.Pagination{
		Page: int(req.Page),
		Size: &req.Size,
		Uid:  userId,
	}
	posts, err := s.biz.ListPost(ctx, pagination)
	if err != nil {
		return &pb.ListPostReply{
			Code: 1,
			Msg:  err.Error(),
		}, err
	}
	pbPosts := make([]*pb.ListPost, len(posts))
	for i, post := range posts {
		pbPosts[i] = &pb.ListPost{
			Id:        post.ID,
			Title:     post.Title,
			Content:   post.Content,
			CreatedAt: post.CreateAt,
			UpdatedAt: post.UpdatedAt,
			UserId:    post.UserID,
			PlateId:   post.Plate.ID,
		}
	}
	return &pb.ListPostReply{
		Code: 0,
		Msg:  "list post success",
		Data: pbPosts,
	}, nil
}
func (s *PostService) ListPubPost(ctx context.Context, req *pb.ListPubPostRequest) (*pb.ListPubPostReply, error) {
	pagination := domain.Pagination{
		Page: int(req.Page),
		Size: &req.Size,
	}
	posts, err := s.biz.ListPubPost(ctx, pagination)
	if err != nil {
		return &pb.ListPubPostReply{
			Code: 1,
			Msg:  err.Error(),
		}, err
	}
	pbPosts := make([]*pb.ListPost, len(posts))
	for i, post := range posts {
		pbPosts[i] = &pb.ListPost{
			Id:        post.ID,
			Title:     post.Title,
			Content:   post.Content,
			CreatedAt: post.CreateAt,
			UpdatedAt: post.UpdatedAt,
			UserId:    post.UserID,
			PlateId:   post.Plate.ID,
		}
	}
	return &pb.ListPubPostReply{
		Code: 0,
		Msg:  "list post success",
		Data: pbPosts,
	}, nil
}

func (s *PostService) ListAdminPost(ctx context.Context, req *pb.ListAdminPostRequest) (*pb.ListAdminPostReply, error) {
	// TODO 暂时保留
	return &pb.ListAdminPostReply{}, nil
}
func (s *PostService) DetailPost(ctx context.Context, req *pb.DetailPostRequest) (*pb.DetailPostReply, error) {
	userId, err := s.getUserId(ctx)
	if err != nil {
		return &pb.DetailPostReply{
			Code: 1,
			Msg:  err.Error(),
		}, err
	}
	post, err := s.biz.GetPost(ctx, req.PostId, userId)
	if err != nil {
		return &pb.DetailPostReply{
			Code: 1,
			Msg:  err.Error(),
		}, err
	}
	return &pb.DetailPostReply{
		Code: 0,
		Msg:  "get post success",
		Data: &pb.DetailPost{
			Id:        post.ID,
			Title:     post.Title,
			Content:   post.Content,
			CreatedAt: post.CreateAt,
			UpdatedAt: post.UpdatedAt,
			UserId:    post.UserID,
			PlateId:   post.Plate.ID,
		},
	}, nil
}
func (s *PostService) DetailPubPost(ctx context.Context, req *pb.DetailPubPostRequest) (*pb.DetailPubPostReply, error) {
	post, err := s.biz.GetPubPost(ctx, req.PostId)
	if err != nil {
		return &pb.DetailPubPostReply{
			Code: 1,
			Msg:  err.Error(),
		}, err
	}
	return &pb.DetailPubPostReply{
		Code: 0,
		Msg:  "get pub post success",
		Data: &pb.DetailPost{
			Id:           post.ID,
			Title:        post.Title,
			Content:      post.Content,
			CreatedAt:    post.CreateAt,
			UpdatedAt:    post.UpdatedAt,
			UserId:       post.UserID,
			PlateId:      post.Plate.ID,
			LikeCount:    post.LikeNum,
			CollectCount: post.CollectNum,
			ViewCount:    post.ViewNum,
		},
	}, nil
}
func (s *PostService) DetailAdminPost(ctx context.Context, req *pb.DetailAdminPostRequest) (*pb.DetailAdminPostReply, error) {
	// TODO 暂时保留
	return &pb.DetailAdminPostReply{}, nil
}
func (s *PostService) GetPostStats(ctx context.Context, req *empty.Empty) (*pb.GetPostStatsReply, error) {
	// TODO 暂时保留
	return &pb.GetPostStatsReply{}, nil
}
func (s *PostService) LikePost(ctx context.Context, req *pb.LikePostRequest) (*pb.LikePostReply, error) {
	// TODO 暂时保留
	return &pb.LikePostReply{}, nil
}
func (s *PostService) CollectPost(ctx context.Context, req *pb.CollectPostRequest) (*pb.CollectPostReply, error) {
	// TODO 暂时保留
	return &pb.CollectPostReply{}, nil
}

func (s *PostService) getUserId(ctx context.Context) (int64, error) {
	// Initialize etcd client
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return -1, err
	}
	defer etcdClient.Close()
	// Initialize etcd registry
	r := etcd.New(etcdClient)
	// Initialize gRPC connection with service discovery
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///linkme-user"),
		grpc.WithDiscovery(r),
	)
	if err != nil {
		return -1, err
	}
	defer conn.Close()
	// Initialize user client
	userClient := userpb.NewUserClient(conn)
	token := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ0NFpHOWJuQ2RUaTlCWGFVZHpoNmhpcFUxQll4S0daZiIsImV4cCI6MTcyMDk1NjgwMCwiVWlkIjoyMTE0NDE5NDIzMTgzMjU3NiwiU3NpZCI6ImY2NDhlYWM2LWZkODItNDY2MS1hNTZlLWZkNjI0MTJiMmNmYiIsIlVzZXJBZ2VudCI6Ik1vemlsbGEvNS4wIChNYWNpbnRvc2g7IEludGVsIE1hYyBPUyBYIDEwXzE1XzcpIEFwcGxlV2ViS2l0LzUzNy4zNiAoS0hUTUwsIGxpa2UgR2Vja28pIENocm9tZS8xMjQuMC4wLjAgU2FmYXJpLzUzNy4zNiIsIkNvbnRlbnRUeXBlIjoiYXBwbGljYXRpb24vanNvbiJ9.cOHuBPXVkbgd4vmY9wlEGkWYpFIz4-9y5as5Wp2gwE-BZE1gjXIxsaoAHDDRxVMUQDpeNUfwSXUXqK_Y01dqFg"
	// Create request to get user info
	req := &userpb.GetUserInfoRequest{Token: token}
	// Call GetUserInfo method
	info, err := userClient.GetUserInfo(ctx, req)
	if err != nil {
		return -1, err
	}

	return info.UserId, nil
}
