package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNamePromStrategy = "prom_strategies"

// PromStrategy mapped from table <prom_strategies>
type PromStrategy struct {
	ID           int32            `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true" json:"id"`
	GroupID      int32            `gorm:"column:group_id;type:int unsigned;not null;comment:所属规则组ID" json:"group_id"`                                             // 所属规则组ID
	Alert        string           `gorm:"column:alert;type:varchar(64);not null;comment:规则名称" json:"alert"`                                                       // 规则名称
	Expr         string           `gorm:"column:expr;type:text;not null;comment:prom ql" json:"expr"`                                                             // prom ql
	For          string           `gorm:"column:for;type:varchar(64);not null;default:10s;comment:持续时间" json:"for"`                                               // 持续时间
	Labels       string           `gorm:"column:labels;type:json;not null;comment:标签" json:"labels"`                                                              // 标签
	Annotations  string           `gorm:"column:annotations;type:json;not null;comment:告警文案" json:"annotations"`                                                  // 告警文案
	AlertLevelID int32            `gorm:"column:alert_level_id;type:int;not null;index:idx__alart_level_id,priority:1;comment:告警等级dict ID" json:"alert_level_id"` // 告警等级dict ID
	Status       int32            `gorm:"column:status;type:tinyint;not null;default:1;comment:启用状态: 1启用;2禁用" json:"status"`                                      // 启用状态: 1启用;2禁用
	CreatedAt    time.Time        `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"`                     // 创建时间
	UpdatedAt    time.Time        `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"`                     // 更新时间
	DeletedAt    gorm.DeletedAt   `gorm:"column:deleted_at;type:timestamp;comment:删除时间" json:"deleted_at"`                                                        // 删除时间
	AlarmPages   []*PromAlarmPage `gorm:"References:ID;foreignKey:ID;joinForeignKey:PromStrategyID;joinReferences:AlarmPageID;many2many:prom_strategy_alarm_pages" json:"alarm_pages"`
	Categories   []*PromDict      `gorm:"References:ID;foreignKey:ID;joinForeignKey:PromStrategyID;joinReferences:DictID;many2many:prom_strategy_categories" json:"categories"`
	AlertLevel   *PromDict        `gorm:"foreignKey:AlertLevelID" json:"alert_level"`
	GroupInfo    *PromGroup       `gorm:"foreignKey:GroupID" json:"group_info"`
}

// TableName PromStrategy's table name
func (*PromStrategy) TableName() string {
	return TableNamePromStrategy
}
