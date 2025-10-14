package contents

import (
	"github.com/BooBooStory/config/models"
	"gorm.io/gorm"
)

type Repository interface {
	Create(content models.StoryContent) (*models.StoryContent, error)
	Update(id uint, content models.StoryContent) (*models.StoryContent, error)
	Delete(id uint) error
	FindAll(story_id uint) ([]models.StoryContent, error)
}

type repository struct {
	DB *gorm.DB
}

// Create implements Repository.
func (r *repository) Create(content models.StoryContent) (*models.StoryContent, error) {
	if err := r.DB.Create(content).Error; err !=nil{
		return nil, err
	}
	return  &content, nil
}

// Delete implements Repository.
func (r *repository) Delete(id uint) error {
	return r.DB.Delete(id).Error
}

// FindAll implements Repository.
func (r *repository) FindAll(story_id uint) ([]models.StoryContent, error) {
	var StoryContent []models.StoryContent
	
	if err := r.DB.Where("story_id = ?", story_id).Find(&StoryContent).Error; err != nil {
		return []models.StoryContent{}, err
	}
	return StoryContent, nil
}

// Update implements Repository.
func (r *repository) Update(id uint, content models.StoryContent) (*models.StoryContent, error) {
	if err := r.DB.Save(content).Error; err != nil{
		return nil, err
	}
	return &content, nil
}

func NewRepository(DB *gorm.DB) Repository {
	return &repository{
		DB: DB,
	}
}
