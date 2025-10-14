package models

import (
	"time"

	"gorm.io/gorm"
)

type Story struct {
	gorm.Model
	Title        string
	Description  string
	CoverImage   string
	Status       string     `gorm:"default:'draft'"` 
	Approved     bool        `gorm:"default:false"`
	AgeLevel     string
	Difficulty   int
	PublishedAt *time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	CreatorID    uint

	Creator *Creator `gorm:"foreignKey:CreatorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	Categories []*Category `gorm:"many2many:story_categories;"`
	Contents   []StoryContent `gorm:"foreignKey:StoryID"`
	RareWords  []StoryContentRareWord `gorm:"foreignKey:StoryID"`
	ProgressReaders []ProgressRead `gorm:"foreignKey:StoryID"`
	Quiz *Quiz `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:StoryID"`

	// 🔖 Relasi bookmark (many-to-many)
	BookmarkedBy []*User `gorm:"many2many:story_bookmarks;"`
}
