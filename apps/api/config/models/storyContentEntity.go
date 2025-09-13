package models

import "time"

// StoryContent merepresentasikan tabel 'story_contents'
type StoryContent struct {
	ID        uint `gorm:"primaryKey"`
	StoryID   uint
	PageOrder int
	Text      string
	ImageUrl  string
	AudioUrl  string
	CreatedAt time.Time
}