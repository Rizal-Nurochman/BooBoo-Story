package models

import "time"

type UserAchievement struct {
	ID            uint `gorm:"primaryKey"`
	UserID        uint
	AchievementID uint
	EarnedAt      time.Time

	User        User        `gorm:"foreignKey:UserID"`
	Achievement Achievement `gorm:"foreignKey:AchievementID"`
}
