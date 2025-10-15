package stories

import (
	"errors"
	"github.com/BooBooStory/config/models"
	"github.com/BooBooStory/config/validations"
	"github.com/BooBooStory/utils"
)

type Service interface {
	Create(story validations.StoryCreate) (*models.Story, error)
	Update(idStr string, story validations.StoryUpdate) (*models.Story, error)
	Delete(idStr string) error
	GetAll(includes []string, page int, limit int, q string) ([]models.Story, int64, error)
	GetByID(idStr string, includes []string) (*models.Story, error)
}

type service struct {
	repo Repository
}

// ✅ CREATE
func (s *service) Create(story validations.StoryCreate) (*models.Story, error) {
	newStory := models.Story{
		Title:       story.Title,
		Description: story.Description,
		CoverImage:  story.CoverImage,
		AgeLevel:    story.AgeLevel,
		Difficulty:  story.Difficulty,
		CreatorID:   story.CreatorID,
	}

	createdStory, err := s.repo.Create(newStory)
	if err != nil {
		return nil, err
	}

	return createdStory, nil
}

// ✅ DELETE
func (s *service) Delete(idStr string) error {
	id, err := utils.ToUint(idStr)
	if err != nil {
		return err
	}

	return s.repo.Delete(uint(id))
}

// ✅ GET ALL
func (s *service) GetAll(includes []string, page int, limit int, q string) ([]models.Story, int64, error) {
	return s.repo.FindAll(includes, page, limit, q)
}

// ✅ GET BY ID
func (s *service) GetByID(idStr string, includes []string) (*models.Story, error) {
	id, err := utils.ToUint(idStr)
	if err != nil {
		return nil, err
	}

	return s.repo.FindByID(uint(id), includes)
}

// ✅ UPDATE
func (s *service) Update(idStr string, story validations.StoryUpdate) (*models.Story, error) {
	id, err := utils.ToUint(idStr)
	if err != nil {
		return nil, err
	}

	existing, err := s.repo.FindByID(uint(id), []string{})
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return nil, errors.New("story not found")
	}

	// Hanya update field yang dikirim user (non-nil)
	if story.Title != nil {
		existing.Title = *story.Title
	}
	if story.Description != nil {
		existing.Description = *story.Description
	}
	if story.CoverImage != nil {
		existing.CoverImage = *story.CoverImage
	}
	if story.AgeLevel != nil {
		existing.AgeLevel = *story.AgeLevel
	}
	if story.Difficulty != nil {
		existing.Difficulty = *story.Difficulty
	}

	updatedStory, err := s.repo.Update(uint(id), *existing)
	if err != nil {
		return nil, err
	}

	return updatedStory, nil
}

func NewService(repo Repository) Service {
	return &service{repo}
}
