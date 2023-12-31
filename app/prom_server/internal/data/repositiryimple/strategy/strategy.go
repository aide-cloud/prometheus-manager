package strategy

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"

	"prometheus-manager/app/prom_server/internal/biz/bo"
	"prometheus-manager/app/prom_server/internal/biz/do"
	"prometheus-manager/app/prom_server/internal/biz/do/basescopes"
	"prometheus-manager/app/prom_server/internal/biz/repository"
	"prometheus-manager/app/prom_server/internal/biz/vo"
	"prometheus-manager/app/prom_server/internal/data"
	"prometheus-manager/pkg/util/slices"
)

var _ repository.StrategyRepo = (*strategyRepoImpl)(nil)

type (
	strategyRepoImpl struct {
		repository.UnimplementedStrategyRepo

		data *data.Data
		log  *log.Helper

		strategyGroupRepo repository.StrategyGroupRepo
	}
)

func (l *strategyRepoImpl) ListStrategyByIds(ctx context.Context, ids []uint32) ([]*bo.StrategyBO, error) {
	modelList := make([]*do.PromStrategy, 0, len(ids))
	if err := l.data.DB().WithContext(ctx).Scopes(basescopes.InIds(ids...)).Find(&modelList).Error; err != nil {
		return nil, err
	}

	list := slices.To(modelList, func(item *do.PromStrategy) *bo.StrategyBO {
		return bo.StrategyModelToBO(item)
	})
	return list, nil
}

// CreateStrategy 创建策略 TODO 需要增加事物, 保证数据一致性
func (l *strategyRepoImpl) CreateStrategy(ctx context.Context, strategyBO *bo.StrategyBO) (*bo.StrategyBO, error) {
	newStrategy := strategyBO.ToModel()
	// 替换报警页面和分类
	alarmPages := slices.To(strategyBO.AlarmPageIds, func(pageId uint32) *do.PromAlarmPage {
		return &do.PromAlarmPage{
			BaseModel: do.BaseModel{ID: pageId},
		}
	})
	categories := slices.To(strategyBO.CategoryIds, func(categoryId uint32) *do.PromDict {
		return &do.PromDict{
			BaseModel: do.BaseModel{ID: categoryId},
		}
	})

	err := l.data.DB().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		txCtx := basescopes.WithTx(ctx, tx)
		if err := tx.WithContext(txCtx).Create(newStrategy).Error; err != nil {
			return err
		}

		if err := tx.WithContext(txCtx).Model(newStrategy).Association(basescopes.PreloadKeyAlarmPages).Replace(alarmPages); err != nil {
			return err
		}
		if err := tx.WithContext(txCtx).Model(newStrategy).Association(basescopes.PreloadKeyCategories).Replace(categories); err != nil {
			return err
		}

		// 更新策略组的策略数量
		if err := l.strategyGroupRepo.UpdateStrategyCount(txCtx, strategyBO.GroupId); err != nil {
			return err
		}

		// 更新策略组的启用策略数量
		if err := l.strategyGroupRepo.UpdateEnableStrategyCount(txCtx, strategyBO.GroupId); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return bo.StrategyModelToBO(newStrategy), nil
}

// UpdateStrategyById 更新策略 TODO 需要增加事物, 保证数据一致性
func (l *strategyRepoImpl) UpdateStrategyById(ctx context.Context, id uint32, strategyBO *bo.StrategyBO) (*bo.StrategyBO, error) {
	detail, err := l.getStrategyById(ctx, id)
	if err != nil {
		return nil, err
	}

	newStrategy := strategyBO.ToModel()
	newStrategy.ID = detail.ID
	newStrategy.Status = detail.Status
	// 替换报警页面和分类
	alarmPages := slices.To(strategyBO.AlarmPageIds, func(pageId uint32) *do.PromAlarmPage {
		return &do.PromAlarmPage{
			BaseModel: do.BaseModel{ID: pageId},
		}
	})
	categories := slices.To(strategyBO.CategoryIds, func(categoryId uint32) *do.PromDict {
		return &do.PromDict{
			BaseModel: do.BaseModel{ID: categoryId},
		}
	})

	newStrategyMap := newStrategy.ToMap()
	err = l.data.DB().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		txCtx := basescopes.WithTx(ctx, tx)
		if err = tx.WithContext(txCtx).Model(detail).Scopes(basescopes.InIds(id)).Updates(newStrategyMap).Error; err != nil {
			return err
		}

		if err = tx.WithContext(txCtx).Model(detail).Association(basescopes.PreloadKeyAlarmPages).Replace(&alarmPages); err != nil {
			return err
		}
		if err = tx.WithContext(txCtx).Model(detail).Association(basescopes.PreloadKeyCategories).Replace(&categories); err != nil {
			return err
		}

		if detail.Status != newStrategy.Status && !newStrategy.Status.IsUnknown() {
			// 更新策略组的启用策略数量
			if err = l.strategyGroupRepo.UpdateEnableStrategyCount(txCtx, strategyBO.GroupId); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return bo.StrategyModelToBO(newStrategy), nil
}

func (l *strategyRepoImpl) BatchUpdateStrategyStatusByIds(ctx context.Context, status vo.Status, ids []uint32) error {
	return l.data.DB().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		txCtx := basescopes.WithTx(ctx, tx)
		if err := tx.WithContext(txCtx).Scopes(basescopes.InIds(ids...)).Updates(&do.PromStrategy{Status: status}).Error; err != nil {
			return err
		}
		// 更新策略组的启用策略数量
		if err := l.strategyGroupRepo.UpdateEnableStrategyCount(txCtx, ids...); err != nil {
			return err
		}
		return nil
	})
}

func (l *strategyRepoImpl) DeleteStrategyByIds(ctx context.Context, ids ...uint32) error {
	var detailList []*do.PromStrategy
	if err := l.data.DB().Scopes(basescopes.InIds(ids...)).Find(&detailList).Error; err != nil {
		return err
	}

	return l.data.DB().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		txCtx := basescopes.WithTx(ctx, tx)
		if err := tx.WithContext(txCtx).Scopes(basescopes.InIds(ids...)).Delete(&do.PromStrategy{}).Error; err != nil {
			return err
		}
		groupIds := slices.To(detailList, func(i *do.PromStrategy) uint32 {
			return i.GroupID
		})
		// 更新策略组的策略数量
		if err := l.strategyGroupRepo.UpdateStrategyCount(txCtx, groupIds...); err != nil {
			return err
		}

		// 更新策略组的启用策略数量
		if err := l.strategyGroupRepo.UpdateEnableStrategyCount(txCtx, groupIds...); err != nil {
			return err
		}
		return nil
	})
}

func (l *strategyRepoImpl) GetStrategyById(ctx context.Context, id uint32, wheres ...basescopes.ScopeMethod) (*bo.StrategyBO, error) {
	firstStrategy, err := l.getStrategyById(ctx, id, wheres...)
	if err != nil {
		return nil, err
	}
	return bo.StrategyModelToBO(firstStrategy), nil
}

func (l *strategyRepoImpl) getStrategyById(ctx context.Context, id uint32, wheres ...basescopes.ScopeMethod) (*do.PromStrategy, error) {
	var first do.PromStrategy
	if err := l.data.DB().WithContext(ctx).Scopes(append(wheres, basescopes.InIds(id))...).First(&first).Error; err != nil {
		return nil, err
	}
	return &first, nil
}

func (l *strategyRepoImpl) ListStrategy(ctx context.Context, pgInfo basescopes.Pagination, scopes ...basescopes.ScopeMethod) ([]*bo.StrategyBO, error) {
	var listStrategy []*do.PromStrategy

	if err := l.data.DB().WithContext(ctx).Scopes(append(scopes, basescopes.Page(pgInfo))...).Find(&listStrategy).Error; err != nil {
		return nil, err
	}
	if pgInfo != nil {
		var total int64
		if err := l.data.DB().WithContext(ctx).Model(&do.PromStrategy{}).Count(&total).Error; err != nil {
			return nil, err
		}
		pgInfo.SetTotal(total)
	}
	list := slices.To(listStrategy, func(i *do.PromStrategy) *bo.StrategyBO {
		return bo.StrategyModelToBO(i)
	})
	return list, nil
}

func NewStrategyRepo(data *data.Data, strategyGroupRepo repository.StrategyGroupRepo, logger log.Logger) repository.StrategyRepo {
	return &strategyRepoImpl{
		data:              data,
		log:               log.NewHelper(logger),
		strategyGroupRepo: strategyGroupRepo,
	}
}
