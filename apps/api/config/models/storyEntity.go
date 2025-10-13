package models

import (
	"time"

	"gorm.io/gorm"
)

type Story struct {
	gorm.Model
	Title       string
	Description string
	Approved    bool
	AgeLevel    string
	Difficulty  int
	CreatedAt   time.Time
	UserID      uint

	// Relasi Belongs To
	User *User `gorm:"foreignKey:UserID"`

	// Relasi Many-to-Many ke Category
	Categories []*Category `gorm:"many2many:story_categories;"`

	// Relasi Has Many
	Contents        []StoryContent         `gorm:"foreignKey:StoryID"`
	RareWords       []StoryContentRareWord `gorm:"foreignKey:StoryID"`
	ProgressReaders []ProgressRead         `gorm:"foreignKey:UserID"`

	// Relasi One To One
	Quiz *Quiz `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:StoryID"`
}
