package biz

import (
	"context"

	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-interactive/domain"
)

type InteractiveBiz struct {
	ir InteractiveRepo
}

type InteractiveRepo interface {
	AddReadCount(ctx context.Context, postId int64, biz string) error
	AddCollectCount(ctx context.Context, postId int64, biz string) error
	AddLikeCount(ctx context.Context, postId int64, biz string) error
	GetInteractive(ctx context.Context, postId int64) (domain.Interactive, error)
	ListInteractive(ctx context.Context, pagination domain.Pagination) ([]domain.Interactive, error)
}

func NewInteractiveBiz(ir InteractiveRepo) *InteractiveBiz {
	return &InteractiveBiz{
		ir: ir,
	}
}

func (ib *InteractiveBiz) AddReadCount(ctx context.Context, postId int64, biz string) error {
	return ib.ir.AddReadCount(ctx, postId, biz)
}

func (ib *InteractiveBiz) AddCollectCount(ctx context.Context, postId int64, biz string) error {
	return ib.ir.AddCollectCount(ctx, postId, biz)
}

func (ib *InteractiveBiz) AddLikeCount(ctx context.Context, postId int64, biz string) error {
	return ib.ir.AddLikeCount(ctx, postId, biz)
}

func (ib *InteractiveBiz) GetInteractive(ctx context.Context, postId int64) (domain.Interactive, error) {
	return ib.ir.GetInteractive(ctx, postId)
}

func (ib *InteractiveBiz) ListInteractive(ctx context.Context, pagination domain.Pagination) ([]domain.Interactive, error) {
	return ib.ir.ListInteractive(ctx, pagination)
}
