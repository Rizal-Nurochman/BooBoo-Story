package models

import (
	"gorm.io/gorm"
)

type StoryBookmark struct {
	gorm.Model
	UserID  uint
	StoryID uint

	User  *User  `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Story *Story `gorm:"foreignKey:StoryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
