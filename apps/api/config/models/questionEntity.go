package models

import "gorm.io/gorm"

// Question merepresentasikan satu pertanyaan dalam sebuah quiz
type Question struct {
	gorm.Model
	QuizID   uint
	Question string

	// Relasi Belongs To
	Quiz Quiz `gorm:"foreignKey:QuizID"`
	// Relasi Has Many
	Answers []Answer `gorm:"foreignKey:QuestionID"`
}