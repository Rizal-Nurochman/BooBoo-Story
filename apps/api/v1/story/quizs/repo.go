package quizs

import (
	"github.com/BooBooStory/config/models"
	"gorm.io/gorm"
)

type quiz struct {
	db *gorm.DB
}

type QuizRepository interface {
	Create(quiz *models.Quiz) (*models.Quiz, error)
	FindAll(includes []string, page int, limit int, q string) ([]models.Quiz, int64, error)
	FindByID(quizID uint) (*models.Quiz, error)
	Delete(quizID uint) error
}

func NewQuizRepository(db *gorm.DB) *quiz {
	return &quiz{db}
}

func (r *quiz) Create(quiz *models.Quiz) (*models.Quiz, error) {
	err := r.db.Create(quiz).Error
	if err != nil {
		return nil, err
	}

	return quiz, nil
}

func (r *quiz) FindAll(includes []string, page int, limit int, q string) ([]models.Quiz, int64, error) {
	var quizzes []models.Quiz
	var total int64
	query := r.db.Model(&models.Quiz{})

	if q != "" {
		query = query.Joins("JOIN questions ON questions.quiz_id = quizzes.id").
			Where("questions.question ILIKE ?", "%"+q+"%")
	}

	countQuery := r.db.Model(&models.Quiz{})
	if q != "" {
		countQuery = countQuery.Joins("JOIN questions ON questions.quiz_id = quizzes.id").
			Where("questions.question ILIKE ?", "%"+q+"%")
	}
	
	err := countQuery.Distinct("quizzes.id").Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	for _, inc := range includes {
		query = query.Preload(inc)
	}

	query = query.Group("quizzes.id")

	offset := (page - 1) * limit
	err = query.Offset(offset).Limit(limit).Find(&quizzes).Error
	if err != nil {
		return nil, 0, err
	}

	return quizzes, total, nil
}

func (r *quiz) FindByID(quizID uint) (*models.Quiz, error) {
	var quiz models.Quiz

	query := r.db.Preload("Questions").Preload("Questions.Answers")

	err := query.First(&quiz, quizID).Error
	if err != nil {
		return nil, err
	}

	return &quiz, nil
}

func (r *quiz) Delete(quizID uint) error {
	err := r.db.Delete(&models.Quiz{}, quizID).Error
	if err != nil {
		return err
	}

	return nil
}