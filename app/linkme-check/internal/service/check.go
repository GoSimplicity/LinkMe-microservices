package service

import (
	"context"
	pb "github.com/GoSimplicity/LinkMe-microservices/api/check/v1"
	postpb "github.com/GoSimplicity/LinkMe-microservices/api/post/v1"
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-check/domain"
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-check/internal/biz"
)

type CheckService struct {
	pb.UnimplementedCheckServer
	postClient postpb.PostClient
	biz        *biz.CheckBiz
}

func NewCheckService(biz *biz.CheckBiz, postClient postpb.PostClient) *CheckService {
	return &CheckService{
		biz:        biz,
		postClient: postClient,
	}
}

func (s *CheckService) CreateCheck(ctx context.Context, req *pb.CreateCheckRequest) (*pb.CreateCheckReply, error) {
	post, err := s.postClient.DetailAdminPost(ctx, &postpb.DetailAdminPostRequest{
		PostId: req.PostId,
	})
	if err != nil {
		return &pb.CreateCheckReply{
			Code: post.Code,
			Msg:  post.Msg,
		}, err
	}
	checkId, err := s.biz.CreateCheck(ctx, domain.Check{
		PostID:  post.Data.Id,
		Title:   post.Data.Title,
		Content: post.Data.Content,
		UserId:  post.Data.UserId,
	})
	if err != nil {
		return &pb.CreateCheckReply{
			Code:    1,
			Msg:     err.Error(),
			CheckId: -1,
		}, err
	}
	return &pb.CreateCheckReply{
		Code:    0,
		Msg:     "create check success",
		CheckId: checkId,
	}, nil
}

func (s *CheckService) DeleteCheck(ctx context.Context, req *pb.DeleteCheckRequest) (*pb.DeleteCheckReply, error) {
	err := s.biz.DeleteCheck(ctx, req.CheckId)
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

func (s *CheckService) UpdateCheck(ctx context.Context, req *pb.UpdateCheckRequest) (*pb.UpdateCheckReply, error) {
	err := s.biz.UpdateCheck(ctx, domain.Check{
		PostID:  req.Check.PostId,
		Title:   req.Check.Title,
		Content: req.Check.Content,
		UserId:  req.Check.UserId,
	})
	if err != nil {
		return &pb.UpdateCheckReply{
			Code: 1,
			Msg:  err.Error(),
		}, err
	}
	return &pb.UpdateCheckReply{
		Code: 0,
		Msg:  "update check success",
	}, nil
}

func (s *CheckService) GetCheckById(ctx context.Context, req *pb.GetCheckByIdRequest) (*pb.GetCheckByIdReply, error) {
	check, err := s.biz.GetCheckById(ctx, req.CheckId)
	if err != nil {
		return &pb.GetCheckByIdReply{
			Code: 1,
			Msg:  err.Error(),
		}, err
	}
	return &pb.GetCheckByIdReply{
		Code: 0,
		Msg:  "get check success",
		Data: &pb.ListOrGetCheck{
			Id:        check.ID,
			PostId:    check.PostID,
			Title:     check.Title,
			Content:   check.Content,
			UserId:    check.UserId,
			Status:    check.Status,
			Remark:    check.Remark,
			CreatedAt: check.CreatedAt,
			UpdatedAt: check.UpdatedAt,
		},
	}, nil
}

func (s *CheckService) ListChecks(ctx context.Context, req *pb.ListChecksRequest) (*pb.ListChecksReply, error) {
	checks, err := s.biz.ListChecks(ctx, domain.Pagination{
		Page: int(req.Pagination.Page),
		Size: &req.Pagination.Size,
		Uid:  req.Pagination.Uid,
	}, &req.Status)
	if err != nil {
		return &pb.ListChecksReply{
			Code: 1,
			Msg:  err.Error(),
		}, err
	}
	pbChecks := make([]*pb.ListOrGetCheck, len(checks))
	for i, check := range checks {
		pbChecks[i] = &pb.ListOrGetCheck{
			Id:        check.ID,
			PostId:    check.PostID,
			Title:     check.Title,
			Content:   check.Content,
			CreatedAt: check.CreatedAt,
			UpdatedAt: check.UpdatedAt,
			UserId:    check.UserId,
			Status:    check.Status,
			Remark:    check.Remark,
		}
	}
	return &pb.ListChecksReply{
		Code: 0,
		Msg:  "list checks success",
		Data: pbChecks,
	}, nil
}

func (s *CheckService) SubmitCheck(ctx context.Context, req *pb.SubmitCheckRequest) (*pb.SubmitCheckReply, error) {
	err := s.biz.SubmitCheck(ctx, req.CheckId, req.Approved)
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

func (s *CheckService) BatchDeleteChecks(ctx context.Context, req *pb.BatchDeleteChecksRequest) (*pb.BatchDeleteChecksReply, error) {
	err := s.biz.BatchDeleteChecks(ctx, req.CheckIds)
	if err != nil {
		return &pb.BatchDeleteChecksReply{
			Code: 1,
			Msg:  err.Error(),
		}, err
	}
	return &pb.BatchDeleteChecksReply{
		Code: 0,
		Msg:  "batch delete checks success",
	}, nil
}

func (s *CheckService) BatchSubmitChecks(ctx context.Context, req *pb.BatchSubmitChecksRequest) (*pb.BatchSubmitChecksReply, error) {
	domainChecks := make([]domain.Check, len(req.Checks))
	for i, check := range req.Checks {
		domainChecks[i] = domain.Check{
			PostID:  check.PostId,
			Title:   check.Title,
			Content: check.Content,
			UserId:  check.UserId,
			Remark:  check.Remark,
		}
	}
	err := s.biz.BatchSubmitChecks(ctx, domainChecks)
	if err != nil {
		return &pb.BatchSubmitChecksReply{
			Code: 1,
			Msg:  err.Error(),
		}, err
	}
	return &pb.BatchSubmitChecksReply{
		Code: 0,
		Msg:  "batch submit checks success",
	}, nil
}
