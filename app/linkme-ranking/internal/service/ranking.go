package service

import (
	"context"

	pb "github.com/GoSimplicity/LinkMe-microservices/api/ranking/v1"
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-ranking/domain"
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-ranking/internal/biz"
)

type RankingService struct {
	pb.UnimplementedRankingServer
	biz *biz.RankingBiz
}

func NewRankingService(biz *biz.RankingBiz) *RankingService {
	return &RankingService{
		biz: biz,
	}
}

func (s *RankingService) TopN(ctx context.Context, req *pb.TopNRequest) (*pb.TopNReply, error) {
	err := s.biz.TopN(ctx)
	if err != nil {
		return &pb.TopNReply{
			Code: 1,
			Msg:  err.Error(),
		}, err
	}
	return &pb.TopNReply{
		Code: 0,
		Msg:  "run topN success",
	}, nil
}
func (s *RankingService) ListTopN(ctx context.Context, req *pb.ListTopNRequest) (*pb.ListTopNReply, error) {
	rankings, err := s.biz.ListTopN(ctx, domain.Pagination{
		Size: &req.Size,
		Page: int(req.Page),
	})
	if err != nil {
		return &pb.ListTopNReply{
			Code: 1,
			Msg:  err.Error(),
		}, err
	}
	rankingSlice := make([]*pb.GetOrListRanking, 0, len(rankings))
	for _, ranking := range rankings {
		rankingSlice = append(rankingSlice, &pb.GetOrListRanking{
			Title:        ranking.Title,
			Content:      ranking.Content,
			UserId:       ranking.UserID,
			PostId:       ranking.ID,
			PlateId:      ranking.Plate.ID,
			CollectCount: ranking.CollectNum,
			ViewCount:    ranking.ViewNum,
			LikeCount:    ranking.LikeNum,
		})
	}
	return &pb.ListTopNReply{
		Code: 0,
		Msg:  "list rankings success",
		Data: rankingSlice,
	}, nil
}
