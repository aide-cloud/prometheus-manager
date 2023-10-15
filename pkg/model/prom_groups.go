package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNamePromGroup = "prom_groups"

// PromGroup mapped from table <prom_groups>
type PromGroup struct {
	ID             int32           `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true" json:"id"`
	Name           string          `gorm:"column:name;type:varchar(64);not null;comment:规则组名称" json:"name"`                                    // 规则组名称
	StrategyCount  int64           `gorm:"column:strategy_count;type:bigint;not null;comment:规则数量" json:"strategy_count"`                      // 规则数量
	Status         int32           `gorm:"column:status;type:tinyint;not null;default:1;comment:启用状态1:启用;2禁用" json:"status"`                   // 启用状态1:启用;2禁用
	Remark         string          `gorm:"column:remark;type:varchar(255);not null;comment:描述信息" json:"remark"`                                // 描述信息
	CreatedAt      time.Time       `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"` // 创建时间
	UpdatedAt      time.Time       `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"` // 更新时间
	DeletedAt      gorm.DeletedAt  `gorm:"column:deleted_at;type:timestamp;comment:删除时间" json:"deleted_at"`                                    // 删除时间
	PromStrategies []*PromStrategy `gorm:"foreignKey:GroupID" json:"prom_strategies"`
	Categories     []*PromDict     `gorm:"References:ID;foreignKey:ID;joinForeignKey:PromGroupID;joinReferences:DictID;many2many:prom_group_categories" json:"categories"`
}

// TableName PromGroup's table name
func (*PromGroup) TableName() string {
	return TableNamePromGroup
}
