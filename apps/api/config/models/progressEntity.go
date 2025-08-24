package models

import "time"

// Progress merepresentasikan progres membaca user
type Progress struct {
	ID           string `gorm:"primaryKey"`
	UserID       string
	StoryID      string
	CurrentPage  int
	Completed    bool
	EarnedPoints int
	LastRead     time.Time

	// Relasi Belongs To
	User  User  `gorm:"foreignKey:UserID"`
	Story Story `gorm:"foreignKey:StoryID"`
}