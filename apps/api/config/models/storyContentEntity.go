package models

type StoryContent struct {
	ID        uint `gorm:"primaryKey"`
	StoryID   uint
	PageOrder int
	Text      string
	ImageUrl  string
	AudioUrl  string
}

