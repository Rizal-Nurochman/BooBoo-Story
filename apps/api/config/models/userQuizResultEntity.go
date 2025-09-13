package models

import "time"

// UserQuizResult merepresentasikan hasil quiz seorang user
type UserQuizResult struct {
	ID          uint `gorm:"primaryKey"`
	UserID      uint
	QuizID      uint
	Score       int
	TimeSpentMs int
	AnsweredAt  time.Time

	// Relasi Belongs To
	User User `gorm:"foreignKey:UserID"`
	Quiz Quiz `gorm:"foreignKey:QuizID"`
}