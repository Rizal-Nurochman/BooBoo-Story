package quizs

import (
	"github.com/BooBooStory/config/models"
	"github.com/BooBooStory/config/validations"
)

type Service interface {
	Create(input validations.CreateQuizInput) (*models.Quiz, error)
	GetAll(includes []string, page int, limit int, q string) ([]models.Quiz, int64, error)
	GetByID(quizID uint) (*models.Quiz, error)
	Delete(quizID uint) error
}

type service struct {
	repository QuizRepository
}

func NewService(repository QuizRepository) *service {
	return &service{repository}
}

func (s *service) Create(input validations.CreateQuizInput) (*models.Quiz, error) {
	quiz := models.Quiz{
		StoryID:    input.StoryID,
		Difficulty: input.Difficulty,
	}

	for _, qInput := range input.Questions {
		question := models.Question{
			Question: qInput.Question,
		}

		for _, aInput := range qInput.Answers {
			question.Answers = append(question.Answers, models.Answer{
				Text:      aInput.Text,
				IsCorrect: aInput.IsCorrect,
			})
		}

		quiz.Questions = append(quiz.Questions, question)
	}

	createdQuiz, err := s.repository.Create(&quiz)
	if err != nil {
		return nil, err
	}

	return createdQuiz, nil
}

func (s *service) GetAll(includes []string, page int, limit int, q string) ([]models.Quiz, int64, error) {
	return s.repository.FindAll(includes, page, limit, q)
}

func (s *service) GetByID(quizID uint) (*models.Quiz, error) {
	return s.repository.FindByID(quizID)
}

func (s *service) Delete(quizID uint) error {
	return s.repository.Delete(quizID)
}