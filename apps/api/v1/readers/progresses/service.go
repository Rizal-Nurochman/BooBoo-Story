package progresses

import (
	"errors"
	"time"

	"github.com/BooBooStory/config/models"
	"github.com/BooBooStory/config/validations"
)

type service struct {
	repository Repository
}

type Service interface {
	Create(userID uint, progress validations.CreateProgressInput) (models.ProgressRead, error)
	GetAll(userID uint, includes []string, page int, limit int, q string) ([]models.ProgressRead, int64, error)
	GetByID(userID uint, progressID uint, includes []string) (*models.ProgressRead, error)
	Update(userID uint, progressID uint, progress validations.UpdateProgressInput) (*models.ProgressRead, error)
	Delete(userID uint, progressID uint) error
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Create(userID uint, input validations.CreateProgressInput) (models.ProgressRead, error) {
	progress := models.ProgressRead{
		UserID:      userID,
		StoryID:     input.StoryID,
		CurrentPage: input.CurrentPage,
		Completed:   false,
		LastRead:    time.Now(),
	}

	return s.repository.Create(progress)
}

func (s *service) GetAll(userID uint, includes []string, page int, limit int, q string) ([]models.ProgressRead, int64, error) {
	progresses, total, err := s.repository.FindAll(userID, includes, page, limit, q)
	if err != nil {
		return nil, 0, err
	}

	return progresses, total, nil
}

func (s *service) GetByID(userID uint, progressID uint, includes []string) (*models.ProgressRead, error) {
	progress, err := s.repository.FindByID(progressID, includes)
	if err != nil {
		return nil, err
	}

	if progress.UserID != userID {
		return nil, nil
	}

	return progress, nil
}

func (s *service) Update(userID uint, progressID uint, input validations.UpdateProgressInput) (*models.ProgressRead, error) {
	progressToUpdate, err := s.repository.FindByID(progressID, nil)
	if err != nil {
		return nil, err
	}

	if progressToUpdate.UserID != userID {
		return nil, errors.New("akses ditolak")
	}

	progressToUpdate.CurrentPage = input.CurrentPage
	progressToUpdate.Completed = input.Completed
	progressToUpdate.LastRead = time.Now()

	return s.repository.Update(progressID, *progressToUpdate)
}

func (s *service) Delete(userID uint, progressID uint) error {
	progress, err := s.repository.FindByID(progressID, nil)
	if err != nil {
		return err 
	}

	if progress.UserID != userID {
		return errors.New("akses ditolak")
	}

	return s.repository.Delete(progressID)
}