package models

type Quiz struct {
	ID         uint `gorm:"primaryKey"`
	StoryID    uint `gorm:"uniqueIndex"`
	Difficulty int

	// Relasi Belongs To
	Story *Story `gorm:"foreignKey:StoryID"` // pakai pointer

	// Relasi Has Many
	Questions       []Question       `gorm:"foreignKey:QuizID"`
	UserQuizResults []UserQuizResult `gorm:"foreignKey:QuizID"`
}
