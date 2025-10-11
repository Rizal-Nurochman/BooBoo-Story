package models

import "gorm.io/gorm"

// StoryContentRareWord merepresentasikan kata-kata langka dalam sebuah cerita
type StoryContentRareWord struct {
	gorm.Model
	Title       string
	Description string
	StoryID     string
}