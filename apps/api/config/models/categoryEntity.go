package models

import "gorm.io/gorm"

// Category merepresentasikan tabel 'categories'
type Category struct {
	gorm.Model
	Name   string
	UserID string

	// Relasi Belongs To
	User User `gorm:"foreignKey:UserID"`
	// Relasi Has Many
	Stories []Story `gorm:"foreignKey:CategoryID"`
}