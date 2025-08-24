package models

import "time"

// StoryContent merepresentasikan tabel 'story_contents'
type StoryContent struct {
	ID        string `gorm:"primaryKey"`
	StoryID   string
	PageOrder int
	Text      string
	ImageUrl  string
	AudioUrl  string
	CreatedAt time.Time
}