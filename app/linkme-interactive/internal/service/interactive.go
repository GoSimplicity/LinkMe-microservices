package service

import (
	"context"

	pb "github.com/GoSimplicity/LinkMe-microservices/api/interactive/v1"
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-interactive/domain"
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-interactive/internal/biz"
)

type InteractiveService struct {
	pb.UnimplementedInteractiveServer
	biz *biz.InteractiveBiz
}

func NewInteractiveService(biz *biz.InteractiveBiz) *InteractiveService {
	return &InteractiveService{
		biz: biz,
	}
}

func (s *InteractiveService) GetInteractive(ctx context.Context, req *pb.GetInteractiveRequest) (*pb.GetInteractiveReply, error) {
	interactive, err := s.biz.GetInteractive(ctx, req.PostId)
	if err != nil {
		return &pb.GetInteractiveReply{
			Code: 1,
			Msg:  err.Error(),
		}, err
	}
	return &pb.GetInteractiveReply{
		Code: 0,
		Msg:  "get interactive success",
		Data: &pb.GetOrListInteractive{
			ID:           interactive.Id,
			BizName:      interactive.BizName,
			CollectCount: interactive.CollectCount,
			LikeCount:    interactive.LikeCount,
			PostId:       interactive.PostId,
			ReadCount:    interactive.ReadCount,
			CreateTime:   interactive.CreateTime,
			UpdateTime:   interactive.UpdateTime,
		},
	}, nil
}
func (s *InteractiveService) ListInteractive(ctx context.Context, req *pb.ListInteractiveRequest) (*pb.ListInteractiveReply, error) {
	interactives, err := s.biz.ListInteractive(ctx, domain.Pagination{
		Page: int(req.Page),
		Size: &req.Size,
	})
	if err != nil {
		return &pb.ListInteractiveReply{
			Code: 1,
			Msg:  err.Error(),
		}, err
	}
	interactiveSlice := make([]*pb.GetOrListInteractive, 0, len(interactives))
	for _, interactive := range interactives {
		interactiveSlice = append(interactiveSlice, &pb.GetOrListInteractive{
			ID:           interactive.Id,
			BizID:        interactive.BizId,
			BizName:      interactive.BizName,
			PostId:       interactive.PostId,
			ReadCount:    interactive.ReadCount,
			LikeCount:    interactive.LikeCount,
			CollectCount: interactive.CollectCount,
			CreateTime:   interactive.CreateTime,
			UpdateTime:   interactive.UpdateTime,
		})
	}
	return &pb.ListInteractiveReply{
		Code: 0,
		Msg:  "list interactives success",
		Data: interactiveSlice,
	}, nil
}
func (s *InteractiveService) AddReadCount(ctx context.Context, req *pb.AddCountRequest) (*pb.AddCountReply, error) {
	err := s.biz.AddReadCount(ctx, req.PostId, req.BizName)
	if err != nil {
		return &pb.AddCountReply{
			Code: 1,
			Msg:  err.Error(),
		}, err
	}
	return &pb.AddCountReply{
		Code: 0,
		Msg:  "add read count success",
	}, nil
}

func (s *InteractiveService) AddLikeCount(ctx context.Context, req *pb.AddCountRequest) (*pb.AddCountReply, error) {
	err := s.biz.AddLikeCount(ctx, req.PostId, req.BizName)
	if err != nil {
		return &pb.AddCountReply{
			Code: 1,
			Msg:  err.Error(),
		}, err
	}
	return &pb.AddCountReply{
		Code: 0,
		Msg:  "add like count success",
	}, nil
}

func (s *InteractiveService) AddCollectCount(ctx context.Context, req *pb.AddCountRequest) (*pb.AddCountReply, error) {
	err := s.biz.AddCollectCount(ctx, req.PostId, req.BizName)
	if err != nil {
		return &pb.AddCountReply{
			Code: 1,
			Msg:  err.Error(),
		}, err
	}
	return &pb.AddCountReply{
		Code: 0,
		Msg:  "add collect count success",
	}, nil
}
