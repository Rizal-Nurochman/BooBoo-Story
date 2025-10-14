package models

import "gorm.io/gorm"

// StoryContentRareWord merepresentasikan kata-kata langka dalam sebuah halaman cerita
type StoryContentRareWord struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Description string `gorm:"type:text"`
	StoryContentID uint `gorm:"not null"`
	StoryContent *StoryContent `gorm:"foreignKey:StoryContentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
