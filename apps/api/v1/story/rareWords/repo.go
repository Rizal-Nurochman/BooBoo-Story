package rarewords

import (
	"github.com/BooBooStory/config/models"
	"gorm.io/gorm"
)

type Repository interface {
	Create(rareWord models.StoryContentRareWord) (*models.StoryContentRareWord, error)
	Update(id uint, rareWord models.StoryContentRareWord) (*models.StoryContentRareWord, error)
	Delete(id uint) error
	FindAll(storyContentID uint) ([]models.StoryContentRareWord, error)
	FindByID(id uint) (*models.StoryContentRareWord, error)
}

type repository struct {
	DB *gorm.DB
}

func (r *repository) Create(rareWord models.StoryContentRareWord) (*models.StoryContentRareWord, error) {
	if err := r.DB.Create(&rareWord).Error; err != nil {
		return nil, err
	}
	return &rareWord, nil
}

func (r *repository) Update(id uint, rareWord models.StoryContentRareWord) (*models.StoryContentRareWord, error) {
	var existing models.StoryContentRareWord
	if err := r.DB.First(&existing, id).Error; err != nil {
		return nil, err
	}
	if err := r.DB.Model(&existing).Updates(rareWord).Error; err != nil {
		return nil, err
	}
	return &existing, nil
}

func (r *repository) Delete(id uint) error {
	return r.DB.Delete(&models.StoryContentRareWord{}, id).Error
}

func (r *repository) FindAll(storyContentID uint) ([]models.StoryContentRareWord, error) {
	var rareWords []models.StoryContentRareWord
	if err := r.DB.Where("story_content_id = ?", storyContentID).Find(&rareWords).Error; err != nil {
		return nil, err
	}
	return rareWords, nil
}

func (r *repository) FindByID(id uint) (*models.StoryContentRareWord, error) {
	var rareWord models.StoryContentRareWord
	if err := r.DB.First(&rareWord, id).Error; err != nil {
		return nil, err
	}
	return &rareWord, nil
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		DB: db,
	}
}
