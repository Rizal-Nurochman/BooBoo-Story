package contents

import (
	"errors"

	"github.com/BooBooStory/config/models"
	"github.com/BooBooStory/config/validations"
	"github.com/BooBooStory/utils"
)

type Service interface {
	Create(content validations.StoryContentCreate) (*models.StoryContent, error)
	Update(idStr string, content validations.StoryContentUpdate) (*models.StoryContent, error)
	Delete(idStr string) error
	GetAll(storyIDStr string) ([]models.StoryContent, error)
}

type service struct {
	repo Repository
}

// ✅ Create: buat konten baru
func (s *service) Create(content validations.StoryContentCreate) (*models.StoryContent, error) {
	newContent := models.StoryContent{
		StoryID:  content.StoryID,
		Text:     content.Text,
		ImageUrl: content.ImageUrl,
		AudioUrl: content.AudioUrl,
	}

	createdContent, err := s.repo.Create(newContent)
	if err != nil {
		return nil, err
	}
	return createdContent, nil
}

// ✅ GetAll: ambil semua konten berdasarkan story_id
func (s *service) GetAll(storyIDStr string) ([]models.StoryContent, error) {
	storyID, err := utils.ToUint(storyIDStr)
	if err != nil {
		return nil, err
	}
	return s.repo.FindAll(storyID)
}



// ✅ Update: ubah konten berdasarkan id
func (s *service) Update(idStr string, content validations.StoryContentUpdate) (*models.StoryContent, error) {
	id, err := utils.ToUint(idStr)
	if err != nil {
		return nil, err
	}

	existing, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return nil, errors.New("content not found")
	}

	if content.Text != nil {
		existing.Text = *content.Text
	}
	if content.ImageUrl != nil {
		existing.ImageUrl = *content.ImageUrl
	}
	if content.AudioUrl != nil {
		existing.AudioUrl = *content.AudioUrl
	}

	updated, err := s.repo.Update(id, *existing)
	if err != nil {
		return nil, err
	}
	return updated, nil
}

// ✅ Delete: hapus konten berdasarkan id
func (s *service) Delete(idStr string) error {
	id, err := utils.ToUint(idStr)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}
