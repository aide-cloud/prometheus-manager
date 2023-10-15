package model

const TableNamePromAlarmPageHistory = "prom_alarm_page_histories"

// PromAlarmPageHistory mapped from table <prom_alarm_page_histories>
type PromAlarmPageHistory struct {
	AlarmPageID int32 `gorm:"column:alarm_page_id;type:int unsigned;primaryKey;uniqueIndex:idx__page_id__history_id,priority:1;comment:报警页面ID" json:"alarm_page_id"` // 报警页面ID
	HistoryID   int32 `gorm:"column:history_id;type:int unsigned;primaryKey;uniqueIndex:idx__page_id__history_id,priority:2;comment:历史ID" json:"history_id"`         // 历史ID
}

// TableName PromAlarmPageHistory's table name
func (*PromAlarmPageHistory) TableName() string {
	return TableNamePromAlarmPageHistory
}
