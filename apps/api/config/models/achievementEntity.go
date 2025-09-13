package models

import "gorm.io/datatypes"

// Achievement merepresentasikan tabel 'achievements'
type Achievement struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
	AssetUrl    string
	Condition   datatypes.JSON

	// Relasi Has Many
	UserAchievements []UserAchievement `gorm:"foreignKey:AchievementID"`
}
