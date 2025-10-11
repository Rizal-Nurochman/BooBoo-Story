package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name     string
	UserID   uint
	ParentID *uint
	Parent   *Category  `gorm:"foreignKey:ParentID"`
	Children []Category `gorm:"foreignKey:ParentID"`
	Stories  []Story    `gorm:"foreignKey:CategoryID"`
}
