package biz

import (
	"context"
	"github.com/GoSimplicity/LinkMe/app/linkme-check/domain"
)

type CheckData interface {
	CreateCheck(ctx context.Context, check domain.Check) (int64, error)                                   // 创建审核记录
	DeleteCheck(ctx context.Context, checkId int64) error                                                 // 删除审核记录
	UpdateCheck(ctx context.Context, check domain.Check) error                                            // 更新审核记录
	GetCheckById(ctx context.Context, checkId int64) (domain.Check, error)                                // 根据ID获取审核记录
	ListChecks(ctx context.Context, pagination domain.Pagination, status *string) ([]domain.Check, error) // 获取审核列表，可按状态过滤
	SubmitCheck(ctx context.Context, checkId int64, approved bool, comments string) error                 // 提交审核，包含通过或拒绝操作
	BatchDeleteChecks(ctx context.Context, checkIds []int64) error                                        // 批量删除审核记录
	BatchSubmitChecks(ctx context.Context, checks []domain.Check) error                                   // 批量提交审核记录
}

type CheckBiz struct {
	CheckData CheckData
}

func NewCheckBiz(CheckData CheckData) *CheckBiz {
	return &CheckBiz{
		CheckData: CheckData,
	}
}

func (cs *CheckBiz) CreateCheck(ctx context.Context, check domain.Check) (int64, error) {
	checkId, err := cs.CheckData.CreateCheck(ctx, check)
	if err != nil {
		return -1, err
	}
	return checkId, nil
}

func (cs *CheckBiz) DeleteCheck(ctx context.Context, checkId int64) error {
	// 实现删除审核记录逻辑
	return nil
}

func (cs *CheckBiz) UpdateCheck(ctx context.Context, check domain.Check) error {
	// 实现更新审核记录逻辑
	return nil
}

func (cs *CheckBiz) GetCheckById(ctx context.Context, checkId int64) (domain.Check, error) {
	// 实现根据ID获取审核记录逻辑
	return domain.Check{}, nil
}

func (cs *CheckBiz) ListChecks(ctx context.Context, pagination domain.Pagination, status *string) ([]domain.Check, error) {
	// 实现获取审核记录列表逻辑，按状态过滤
	return nil, nil
}

func (cs *CheckBiz) SubmitCheck(ctx context.Context, checkId int64, approved bool, comments string) error {
	// 实现提交审核逻辑，包含通过或拒绝操作
	return nil
}

func (cs *CheckBiz) BatchDeleteChecks(ctx context.Context, checkIds []int64) error {
	// 实现批量删除审核记录逻辑
	return nil
}

func (cs *CheckBiz) BatchSubmitChecks(ctx context.Context, checks []domain.Check) error {
	// 实现批量提交审核记录逻辑
	return nil
}
