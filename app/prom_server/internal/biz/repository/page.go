package repository

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"prometheus-manager/app/prom_server/internal/biz/do/basescopes"
	"prometheus-manager/app/prom_server/internal/biz/vo"

	"prometheus-manager/app/prom_server/internal/biz/bo"
)

var _ PageRepo = (*UnimplementedPageRepo)(nil)

type (
	PageRepo interface {
		mustEmbedUnimplemented()
		// CreatePage 创建页面
		CreatePage(ctx context.Context, pageBO *bo.AlarmPageBO) (*bo.AlarmPageBO, error)
		// UpdatePageById 通过id更新页面
		UpdatePageById(ctx context.Context, id uint32, pageBO *bo.AlarmPageBO) (*bo.AlarmPageBO, error)
		// BatchUpdatePageStatusByIds 通过id批量更新页面状态
		BatchUpdatePageStatusByIds(ctx context.Context, status vo.Status, ids []uint32) error
		// DeletePageByIds 通过id删除页面
		DeletePageByIds(ctx context.Context, id ...uint32) error
		// GetPageById 通过id获取页面详情
		GetPageById(ctx context.Context, id uint32) (*bo.AlarmPageBO, error)
		// ListPage 获取页面列表
		ListPage(ctx context.Context, pgInfo basescopes.Pagination, scopes ...basescopes.ScopeMethod) ([]*bo.AlarmPageBO, error)
	}

	UnimplementedPageRepo struct{}
)

func (UnimplementedPageRepo) mustEmbedUnimplemented() {}

func (UnimplementedPageRepo) CreatePage(_ context.Context, _ *bo.AlarmPageBO) (*bo.AlarmPageBO, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePage not implemented")
}

func (UnimplementedPageRepo) UpdatePageById(_ context.Context, _ uint32, _ *bo.AlarmPageBO) (*bo.AlarmPageBO, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePageById not implemented")
}

func (UnimplementedPageRepo) BatchUpdatePageStatusByIds(_ context.Context, _ vo.Status, _ []uint32) error {
	return status.Errorf(codes.Unimplemented, "method BatchUpdatePageStatusByIds not implemented")
}

func (UnimplementedPageRepo) DeletePageByIds(_ context.Context, _ ...uint32) error {
	return status.Errorf(codes.Unimplemented, "method DeletePageByIds not implemented")
}

func (UnimplementedPageRepo) GetPageById(_ context.Context, _ uint32) (*bo.AlarmPageBO, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPageById not implemented")
}

func (UnimplementedPageRepo) ListPage(_ context.Context, _ basescopes.Pagination, _ ...basescopes.ScopeMethod) ([]*bo.AlarmPageBO, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPage not implemented")
}
