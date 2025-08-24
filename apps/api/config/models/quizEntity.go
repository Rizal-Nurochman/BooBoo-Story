package models

// Quiz merepresentasikan tabel 'quizzes'
type Quiz struct {
	ID         string `gorm:"primaryKey"`
	StoryID    string
	Difficulty int

	// Relasi Belongs To
	Story Story `gorm:"foreignKey:StoryID"`
	// Relasi Has Many
	Questions       []Question       `gorm:"foreignKey:QuizID"`
	UserQuizResults []UserQuizResult `gorm:"foreignKey:QuizID"`
}