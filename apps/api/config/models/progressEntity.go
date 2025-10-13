package models

import "time"

// Progress merepresentasikan progres membaca user
type ProgressRead struct {
	ID           uint `gorm:"primaryKey"`
	UserID       uint
	StoryID      uint
	CurrentPage  int
	Completed    bool
	EarnedPoints int
	LastRead     time.Time

	// Relasi Belongs To
	User  User  `gorm:"foreignKey:UserID"`
	Story Story `gorm:"foreignKey:StoryID"`
}