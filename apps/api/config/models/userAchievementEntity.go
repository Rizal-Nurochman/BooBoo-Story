package models

import "time"

// UserAchievement adalah join table antara User dan Achievement
type UserAchievement struct {
	ID            string `gorm:"primaryKey"`
	UserID        string
	AchievementID string
	EarnedAt      time.Time

	// Relasi Belongs To
	User        User        `gorm:"foreignKey:UserID"`
	Achievement Achievement `gorm:"foreignKey:AchievementID"`
}