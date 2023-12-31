package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"prometheus-manager/api/perrors"
	"prometheus-manager/app/prom_server/internal/biz/bo"
	"prometheus-manager/app/prom_server/internal/biz/do/basescopes"
	"prometheus-manager/app/prom_server/internal/biz/repository"
)

type (
	ChatGroupBiz struct {
		log *log.Helper

		chatGroupRepo repository.ChatGroupRepo
	}
)

// NewChatGroupBiz .
func NewChatGroupBiz(chatGroupRepo repository.ChatGroupRepo, logger log.Logger) *ChatGroupBiz {
	return &ChatGroupBiz{
		log:           log.NewHelper(logger),
		chatGroupRepo: chatGroupRepo,
	}
}

// CreateChatGroup 创建通知群机器人hook
func (b *ChatGroupBiz) CreateChatGroup(ctx context.Context, chatGroup *bo.ChatGroupBO) (*bo.ChatGroupBO, error) {
	if chatGroup == nil {
		return nil, perrors.ErrorInvalidParams("参数错误")
	}
	return b.chatGroupRepo.Create(ctx, chatGroup)
}

// GetChatGroupById  获取通知群机器人hook
func (b *ChatGroupBiz) GetChatGroupById(ctx context.Context, id uint32) (*bo.ChatGroupBO, error) {
	return b.chatGroupRepo.Get(ctx, basescopes.InIds(id))
}

// ListChatGroup 获取通知群机器人hook列表
func (b *ChatGroupBiz) ListChatGroup(ctx context.Context, pgInfo basescopes.Pagination, scopes ...basescopes.ScopeMethod) ([]*bo.ChatGroupBO, error) {
	return b.chatGroupRepo.List(ctx, pgInfo, scopes...)
}

// UpdateChatGroupById 更新通知群机器人hook
func (b *ChatGroupBiz) UpdateChatGroupById(ctx context.Context, chatGroup *bo.ChatGroupBO, id uint32) error {
	return b.chatGroupRepo.Update(ctx, chatGroup, basescopes.InIds(id))
}

// DeleteChatGroupById 删除通知群机器人hook
func (b *ChatGroupBiz) DeleteChatGroupById(ctx context.Context, id uint32) error {
	return b.chatGroupRepo.Delete(ctx, basescopes.InIds(id))
}
