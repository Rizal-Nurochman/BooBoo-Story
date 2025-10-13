package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name     string
	UserID   uint

	// Hierarki Category
	ParentID *uint
	Parent   *Category   `gorm:"foreignKey:ParentID"`
	Children []Category  `gorm:"foreignKey:ParentID"`

	// Relasi Many-to-Many ke Story
	Stories []*Story `gorm:"many2many:story_categories;"`
}
