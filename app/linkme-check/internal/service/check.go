package service

import (
	"context"
	"errors"
	"fmt"
	pb "github.com/GoSimplicity/LinkMe-microservices/api/check/v1"
	userpb "github.com/GoSimplicity/LinkMe-microservices/api/user/v1"
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-check/internal/biz"
	"github.com/go-kratos/kratos/v2/transport"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strings"
	"time"
)

type CheckService struct {
	pb.UnimplementedCheckServer
	userClient userpb.UserClient
	biz        *biz.CheckBiz
}

func NewCheckService(userClient userpb.UserClient, biz *biz.CheckBiz) *CheckService {
	return &CheckService{
		userClient: userClient,
		biz:        biz,
	}
}

func (s *CheckService) CreateCheck(ctx context.Context, req *pb.CreateCheckRequest) (*pb.CreateCheckReply, error) {
	err := s.biz.CreateCheck(ctx, biz.Check{
		PostID:    req.PostId,
		Title:     req.Title,
		Content:   req.Content,
		UserID:    req.UserId,
		CreatedAt: time.Now(),
		Status:    biz.UnderReview,
	})
	if err != nil {
		return &pb.CreateCheckReply{
			Code: 1,
			Msg:  err.Error(),
		}, err
	}

	return &pb.CreateCheckReply{
		Code: 0,
		Msg:  "success",
	}, nil
}

func (s *CheckService) DeleteCheck(ctx context.Context, req *pb.DeleteCheckRequest) (*pb.DeleteCheckReply, error) {
	userId, err := s.getUserId(ctx)
	if err != nil {
		return &pb.DeleteCheckReply{
			Code: 1,
			Msg:  err.Error(),
		}, err
	}

	err = s.biz.DeleteCheck(ctx, req.CheckId, userId)
	if err != nil {
		return &pb.DeleteCheckReply{
			Code: 1,
			Msg:  err.Error(),
		}, err
	}

	return &pb.DeleteCheckReply{
		Code: 0,
		Msg:  "delete check success",
	}, nil
}

func (s *CheckService) GetCheckById(ctx context.Context, req *pb.GetCheckByIdRequest) (*pb.GetCheckByIdReply, error) {
	check, err := s.biz.GetCheck(ctx, req.CheckId)
	if err != nil {
		return &pb.GetCheckByIdReply{
			Code: 1,
			Msg:  err.Error(),
		}, err
	}

	return &pb.GetCheckByIdReply{
		Code: 0,
		Data: &pb.ListOrGetCheck{
			Id:        check.ID,
			UserId:    check.UserID,
			PostId:    check.PostID,
			Title:     check.Title,
			Content:   check.Content,
			Status:    uint32(check.Status),
			Remark:    check.Remark,
			CreatedAt: timestamppb.New(check.CreatedAt),
			UpdatedAt: timestamppb.New(check.UpdatedAt),
		},
	}, nil
}

func (s *CheckService) ListChecks(ctx context.Context, req *pb.ListChecksRequest) (*pb.ListChecksReply, error) {
	userId, err := s.getUserId(ctx)
	if err != nil {
		return &pb.ListChecksReply{
			Code: 1,
			Msg:  err.Error(),
		}, err
	}
	fmt.Println(userId)

	checks, err := s.biz.ListChecks(ctx, biz.Pagination{
		Page: int(req.Page),
		Size: &req.Size,
	})
	if err != nil {
		return &pb.ListChecksReply{
			Code: 1,
			Msg:  err.Error(),
		}, err
	}

	listChecks := make([]*pb.ListOrGetCheck, 0, len(checks))

	for _, check := range checks {
		listChecks = append(listChecks, &pb.ListOrGetCheck{
			Id:        check.ID,
			UserId:    check.UserID,
			PostId:    check.PostID,
			Title:     check.Title,
			Content:   check.Content,
			Status:    uint32(check.Status),
			Remark:    check.Remark,
			CreatedAt: timestamppb.New(check.CreatedAt),
			UpdatedAt: timestamppb.New(check.UpdatedAt),
		})
	}

	return &pb.ListChecksReply{
		Code: 0,
		Msg:  "success",
		Data: listChecks,
	}, nil
}

func (s *CheckService) SubmitCheck(ctx context.Context, req *pb.SubmitCheckRequest) (*pb.SubmitCheckReply, error) {
	err := s.biz.UpdateStatus(ctx, req.CheckId, req.Remark, uint8(req.Status))
	if err != nil {
		return &pb.SubmitCheckReply{
			Code: 1,
			Msg:  err.Error(),
		}, err
	}

	return &pb.SubmitCheckReply{
		Code: 0,
		Msg:  "submit check success",
	}, nil
}

// 通过grpc调用linkme-user模块方法，获取userId
func (s *CheckService) getUserId(ctx context.Context) (int64, error) {
	// 从 Kratos 上下文中获取传输信息
	tr, ok := transport.FromServerContext(ctx)
	if !ok {
		return -1, errors.New("failed to get transport from context")
	}

	// 打印所有请求头
	fmt.Println("Request Headers:")
	for k, v := range tr.RequestHeader().Keys() {
		fmt.Printf("%s: %s\n", k, tr.RequestHeader().Get(v))
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
