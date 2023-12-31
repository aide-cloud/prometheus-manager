package basescopes

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	RealtimeAssociationReplaceIntervenes = "AlarmIntervenes"
	RealtimeAssociationUpgradeInfo       = "AlarmUpgradeInfo"
	RealtimeAssociationSuppressInfo      = "AlarmSuppressInfo"
	RealtimeAssociationBeenNotifyMembers = "BeenNotifyMembers"
	RealtimeAssociationBeenChatGroups    = "BeenChatGroups"
	RealtimeAssociationLevel             = "Level"
)

const (
	RealtimeTableFieldEventAt   Field = "event_at"
	RealtimeTableFieldHistoryId Field = "history_id"
	RealtimeTableFieldNote      Field = "note"
	RealtimeTableFieldInstance  Field = "instance"
)

// RealtimeLike 查询关键字
func RealtimeLike(keyword string) ScopeMethod {
	return WhereLikeKeyword(keyword, RealtimeTableFieldNote, RealtimeTableFieldInstance)
}

// RealtimeEventAtDesc 事件时间倒序
func RealtimeEventAtDesc() ScopeMethod {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(RealtimeTableFieldEventAt.Format(DESC).String())
	}
}

// InHistoryIds 查询历史ID列表
func InHistoryIds(historyIds ...uint32) ScopeMethod {
	return WhereInColumn(RealtimeTableFieldHistoryId, historyIds...)
}

// ClauseOnConflict 冲突处理
func ClauseOnConflict() ScopeMethod {
	return func(db *gorm.DB) *gorm.DB {
		return db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: BaseFieldID.String()}, {Name: RealtimeTableFieldHistoryId.String()}},
			DoUpdates: clause.AssignmentColumns([]string{BaseFieldStatus.String()}),
		})
	}
}

// PreloadLevel 预加载级别
func PreloadLevel() ScopeMethod {
	return func(db *gorm.DB) *gorm.DB {
		return db.Preload(RealtimeAssociationLevel)
	}
}
