package service

import (
	"context"
	"linkme-check/domain"
	"linkme-check/internal/biz"

	pb "linkme-check/api/check/v1"
)

type CheckService struct {
	pb.UnimplementedCheckServer
	biz *biz.CheckBiz
}

func NewCheckService(biz *biz.CheckBiz) *CheckService {
	return &CheckService{
		biz: biz,
	}
}

func (s *CheckService) CreateCheck(ctx context.Context, req *pb.CreateCheckRequest) (*pb.CreateCheckReply, error) {
	userId, err := s.biz.CreateCheck(ctx, domain.Check{
		PostID:  req.Check.PostId,
		Title:   req.Check.Title,
		Content: req.Check.Content,
		UserID:  req.Check.UserId,
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
		CheckId: userId,
	}, nil
}

func (s *CheckService) DeleteCheck(ctx context.Context, req *pb.DeleteCheckRequest) (*pb.DeleteCheckReply, error) {
	return &pb.DeleteCheckReply{}, nil
}

func (s *CheckService) UpdateCheck(ctx context.Context, req *pb.UpdateCheckRequest) (*pb.UpdateCheckReply, error) {
	return &pb.UpdateCheckReply{}, nil
}

func (s *CheckService) GetCheckById(ctx context.Context, req *pb.GetCheckByIdRequest) (*pb.GetCheckByIdReply, error) {
	return &pb.GetCheckByIdReply{}, nil
}

func (s *CheckService) ListChecks(ctx context.Context, req *pb.ListChecksRequest) (*pb.ListChecksReply, error) {
	return &pb.ListChecksReply{}, nil
}

func (s *CheckService) SubmitCheck(ctx context.Context, req *pb.SubmitCheckRequest) (*pb.SubmitCheckReply, error) {
	return &pb.SubmitCheckReply{}, nil
}

func (s *CheckService) BatchDeleteChecks(ctx context.Context, req *pb.BatchDeleteChecksRequest) (*pb.BatchDeleteChecksReply, error) {
	return &pb.BatchDeleteChecksReply{}, nil
}

func (s *CheckService) BatchSubmitChecks(ctx context.Context, req *pb.BatchSubmitChecksRequest) (*pb.BatchSubmitChecksReply, error) {
	return &pb.BatchSubmitChecksReply{}, nil
}
