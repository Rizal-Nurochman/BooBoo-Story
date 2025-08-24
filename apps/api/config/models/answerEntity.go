package models

import "gorm.io/gorm"

// Answer merepresentasikan satu pilihan jawaban untuk sebuah pertanyaan
type Answer struct {
	gorm.Model
	Text       string `gorm:"column:answer_text"` // Diganti dari Answer
	IsCorrect  bool   `gorm:"default:false"`
	QuestionID uint

	// Relasi Belongs To
	Question Question `gorm:"foreignKey:QuestionID"`
}