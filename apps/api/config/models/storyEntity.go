package models

import "time"

// Story merepresentasikan tabel 'stories'
type Story struct {
	ID          string `gorm:"primaryKey"`
	Title       string
	Description string
	Approved    bool
	AgeLevel    string
	Difficulty  int
	CreatedAt   time.Time
	UserID      string // Diganti dari AuthorID untuk konsistensi
	CategoryID  uint

	// Relasi Belongs To
	User     User     `gorm:"foreignKey:UserID"`
	Category Category `gorm:"foreignKey:CategoryID"`
	// Relasi Has Many
	Contents   []StoryContent         `gorm:"foreignKey:StoryID"`
	RareWords  []StoryContentRareWord `gorm:"foreignKey:StoryID"`
	Progresses []Progress             `gorm:"foreignKey:StoryID"`
	Quizzes    []Quiz                 `gorm:"foreignKey:StoryID"`
}