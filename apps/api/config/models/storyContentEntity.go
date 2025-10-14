package models

import (
	"time"
)

type StoryContent struct {
	ID        uint `gorm:"primaryKey"`
	StoryID   uint
	PageOrder int
	Text      string
	ImageUrl  string
	AudioUrl  string
	CreatedAt time.Time
}

