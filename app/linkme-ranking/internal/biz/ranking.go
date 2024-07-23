package biz

import (
	"context"

	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-ranking/domain"
)

type RankingBiz struct {
	data RankingData
}

type RankingData interface {
	TopN(ctx context.Context) error
	ListTopN(ctx context.Context, pagination domain.Pagination) ([]domain.RankingPost, error)
}

func NewRankingBiz(data RankingData) *RankingBiz {
	return &RankingBiz{
		data: data,
	}
}

func (b *RankingBiz) TopN(ctx context.Context) error {
	return b.data.TopN(ctx)
}

func (b *RankingBiz) ListTopN(ctx context.Context, pagination domain.Pagination) ([]domain.RankingPost, error) {
	offset := int64(pagination.Page-1) * *pagination.Size
	pagination.Offset = &offset
	return b.data.ListTopN(ctx, pagination)
}
