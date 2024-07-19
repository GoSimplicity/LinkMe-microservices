package biz

import (
	"context"
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-check/domain"
)

type CheckData interface {
	CreateCheck(ctx context.Context, check domain.Check) (int64, error)
	DeleteCheck(ctx context.Context, checkId int64) error
	UpdateCheck(ctx context.Context, check domain.Check) error
	GetCheckById(ctx context.Context, checkId int64) (domain.Check, error)
	ListChecks(ctx context.Context, pagination domain.Pagination, status *string) ([]domain.Check, error)
	SubmitCheck(ctx context.Context, checkId int64, approved bool) error
	BatchDeleteChecks(ctx context.Context, checkIds []int64) error
	BatchSubmitChecks(ctx context.Context, checks []domain.Check) error
}

type CheckBiz struct {
	CheckData CheckData
}

func NewCheckBiz(CheckData CheckData) *CheckBiz {
	return &CheckBiz{CheckData: CheckData}
}

func (cs *CheckBiz) CreateCheck(ctx context.Context, check domain.Check) (int64, error) {

	return cs.CheckData.CreateCheck(ctx, check)
}

func (cs *CheckBiz) DeleteCheck(ctx context.Context, checkId int64) error {
	return cs.CheckData.DeleteCheck(ctx, checkId)
}

func (cs *CheckBiz) UpdateCheck(ctx context.Context, check domain.Check) error {
	return cs.CheckData.UpdateCheck(ctx, check)
}

func (cs *CheckBiz) GetCheckById(ctx context.Context, checkId int64) (domain.Check, error) {
	check, err := cs.CheckData.GetCheckById(ctx, checkId)
	if err != nil {
		return domain.Check{}, err
	}
	return check, nil
}

func (cs *CheckBiz) ListChecks(ctx context.Context, pagination domain.Pagination, status *string) ([]domain.Check, error) {
	offset := int64(pagination.Page-1) * *pagination.Size
	pagination.Offset = &offset
	if *status == "" {
		return cs.CheckData.ListChecks(ctx, pagination, nil)
	}
	return cs.CheckData.ListChecks(ctx, pagination, status)
}

func (cs *CheckBiz) SubmitCheck(ctx context.Context, checkId int64, approved bool) error {
	return cs.CheckData.SubmitCheck(ctx, checkId, approved)
}

func (cs *CheckBiz) BatchDeleteChecks(ctx context.Context, checkIds []int64) error {
	return cs.CheckData.BatchDeleteChecks(ctx, checkIds)
}

func (cs *CheckBiz) BatchSubmitChecks(ctx context.Context, checks []domain.Check) error {
	return cs.CheckData.BatchSubmitChecks(ctx, checks)
}
