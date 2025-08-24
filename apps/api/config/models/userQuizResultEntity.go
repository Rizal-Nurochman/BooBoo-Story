package models

import "time"

// UserQuizResult merepresentasikan hasil quiz seorang user
type UserQuizResult struct {
	ID          string `gorm:"primaryKey"`
	UserID      string
	QuizID      string
	Score       int
	TimeSpentMs int
	AnsweredAt  time.Time

	// Relasi Belongs To
	User User `gorm:"foreignKey:UserID"`
	Quiz Quiz `gorm:"foreignKey:QuizID"`
}