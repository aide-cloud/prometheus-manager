package strategy

import (
	query "github.com/aide-cloud/gorm-normalize"
	"gorm.io/gorm"
)

// LikeStrategy 策略
func LikeStrategy(keyword string) query.ScopeMethod {
	return func(db *gorm.DB) *gorm.DB {
		if keyword == "" {
			return db
		}
		return db.Where("name LIKE?", keyword+"%")
	}
}

// StatusEQ 状态
func StatusEQ(status int32) query.ScopeMethod {
	return func(db *gorm.DB) *gorm.DB {
		if status == 0 {
			return db
		}
		return db.Where("status = ?", status)
	}
}