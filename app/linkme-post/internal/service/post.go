package service

import (
	"context"
	"errors"
	pb "github.com/GoSimplicity/LinkMe-microservices/api/post/v1"
	userpb "github.com/GoSimplicity/LinkMe-microservices/api/user/v1"
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-post/internal/biz"
	"github.com/go-kratos/kratos/v2/transport"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strings"
	"time"
)

type PostService struct {
	pb.UnimplementedPostServer
	userClient userpb.UserClient
	biz        *biz.PostBiz
}

func NewPostService(biz *biz.PostBiz, userClient userpb.UserClient) *PostService {
	return &PostService{
		biz:        biz,
		userClient: userClient,
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

	post, err := s.biz.CreatePost(ctx, biz.Post{
		AuthorID: userId,
		Content:  req.Content,
		PlateID:  req.PlateId,
		Title:    req.Title,
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
		Data: post,
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

	err = s.biz.UpdatePost(ctx, biz.Post{
		ID:       req.PostId,
		AuthorID: userId,
		Content:  req.Content,
		PlateID:  req.PlateId,
		Title:    req.Title,
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

func (s *PostService) UpdatePostStatus(ctx context.Context, req *pb.UpdatePostStatusRequest) (*pb.UpdatePostStatusReply, error) {
	userId, err := s.getUserId(ctx)
	if err != nil {
		return &pb.UpdatePostStatusReply{
			Code: 1,
			Msg:  err.Error(),
		}, err
	}

	err = s.biz.UpdatePostStatus(ctx, biz.Post{
		ID:       req.PostId,
		AuthorID: userId,
		Status:   uint8(req.Status),
	})
	if err != nil {
		return &pb.UpdatePostStatusReply{
			Code: 1,
			Msg:  err.Error(),
		}, err
	}

	return &pb.UpdatePostStatusReply{
		Code: 0,
		Msg:  "update post status success",
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

	err = s.biz.DeletePost(ctx, biz.Post{
		ID:       req.PostId,
		AuthorID: userId,
	})
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
	userId, err := s.getUserId(ctx)
	if err != nil {
		return &pb.PublishPostReply{
			Code: 1,
			Msg:  err.Error(),
		}, err
	}

	err = s.biz.PublishPost(ctx, biz.Post{
		ID:       req.PostId,
		AuthorID: userId,
	})
	if err != nil {
		return &pb.PublishPostReply{
			Code: 1,
			Msg:  err.Error(),
		}, err
	}

	return &pb.PublishPostReply{
		Code: 0,
		Msg:  "publish post success",
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

	posts, err := s.biz.ListPost(ctx, biz.Pagination{
		Page: int(req.Page),
		Size: &req.Size,
		Uid:  userId,
	})
	if err != nil {
		return &pb.ListPostReply{
			Code: 1,
			Msg:  err.Error(),
		}, err
	}

	// 创建切片来保存结果
	listPosts := make([]*pb.ListPost, len(posts))

	for i, post := range posts {
		listPosts[i] = &pb.ListPost{
			Id:        post.ID,
			Title:     post.Title,
			Content:   post.Content,
			UserId:    post.AuthorID,
			PlateId:   post.PlateID,
			CreatedAt: timestamppb.New(post.CreatedAt),
			UpdatedAt: timestamppb.New(post.UpdatedAt),
		}
	}

	return &pb.ListPostReply{
		Code: 0,
		Msg:  "list post success",
		Data: listPosts,
	}, nil
}

func (s *PostService) ListPubPost(ctx context.Context, req *pb.ListPubPostRequest) (*pb.ListPubPostReply, error) {
	userId, err := s.getUserId(ctx)
	if err != nil {
		return &pb.ListPubPostReply{
			Code: 1,
			Msg:  "Failed to retrieve user ID: " + err.Error(),
		}, err
	}

	// 调用业务逻辑层获取帖子列表
	posts, err := s.biz.ListPubPost(ctx, biz.Pagination{
		Page: int(req.Page),
		Size: &req.Size,
		Uid:  userId,
	})
	if err != nil {
		return &pb.ListPubPostReply{
			Code: 1,
			Msg:  "Failed to list published posts: " + err.Error(),
		}, err
	}

	// 预先分配切片，避免在循环中动态分配内存
	listPubPosts := make([]*pb.ListPost, len(posts))

	// 使用 for range 遍历 posts 切片
	for i, post := range posts {
		listPubPosts[i] = &pb.ListPost{
			Id:        post.ID,
			Title:     post.Title,
			Content:   post.Content,
			UserId:    post.AuthorID,
			PlateId:   post.PlateID,
			CreatedAt: timestamppb.New(post.CreatedAt),
			UpdatedAt: timestamppb.New(post.UpdatedAt),
		}
	}

	// 返回成功响应
	return &pb.ListPubPostReply{
		Code: 0,
		Msg:  "List published posts success",
		Data: listPubPosts,
	}, nil
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
		Msg:  "detail post success",
		Data: &pb.DetailPost{
			Id:        post.ID,
			Title:     post.Title,
			Content:   post.Content,
			UserId:    post.AuthorID,
			PlateId:   post.PlateID,
			CreatedAt: timestamppb.New(post.CreatedAt),
			UpdatedAt: timestamppb.New(post.UpdatedAt),
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
		Msg:  "detail post success",
		Data: &pb.DetailPost{
			Id:        post.ID,
			Title:     post.Title,
			Content:   post.Content,
			UserId:    post.AuthorID,
			PlateId:   post.PlateID,
			CreatedAt: timestamppb.New(post.CreatedAt),
			UpdatedAt: timestamppb.New(post.UpdatedAt),
		},
	}, nil
}

// 通过grpc调用linkme-user模块方法，获取userId
func (s *PostService) getUserId(ctx context.Context) (int64, error) {
	// 从 Kratos 上下文中获取传输信息
	tr, ok := transport.FromServerContext(ctx)
	if !ok {
		return -1, errors.New("failed to get transport from context")
	}

	// 获取 Authorization 头
	token := tr.RequestHeader().Get("Authorization")
	if token == "" {
		return -1, errors.New("authorization token not provided")
	}

	// 移除 "Bearer " 前缀
	tokenStr := strings.TrimPrefix(token, "Bearer ")
	if tokenStr == "" {
		return -1, errors.New("authorization token is empty after trim")
	}

	// 为 userClient.GetUserInfo 设置超时时间
	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel() // 确保在函数退出时取消上下文，释放资源

	// 调用 userClient 获取用户信息
	info, err := s.userClient.GetUserInfo(timeoutCtx, &userpb.GetUserInfoRequest{Token: tokenStr})
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			// 如果超时，返回具体的超时错误信息
			return -1, errors.New("getUserInfo request timed out")
		}
		return -1, err
	}

	return info.UserId, nil
}
