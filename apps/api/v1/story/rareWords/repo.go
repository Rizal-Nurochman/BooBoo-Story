package rarewords

import (
	"github.com/BooBooStory/config/models"
	"gorm.io/gorm"
)

type Repository interface {
	Create(rareWord models.StoryContentRareWord) (*models.StoryContentRareWord, error)
	Update(id uint, rareWord models.StoryContentRareWord) (*models.StoryContentRareWord, error)
	Delete(id uint) error
	FindAll(story_id uint) ([]models.StoryContentRareWord, error)
}

type repository struct {
	DB *gorm.DB
}


func (r *repository) Create(rareWord models.StoryContentRareWord) (*models.StoryContentRareWord, error) {
	if err := r.DB.Create(rareWord).Error; err !=nil{
		return nil, err
	}	

	return  &rareWord, nil
}

// Delete implements Repository.
func (r *repository) Delete(id uint) error {
	return  r.DB.Delete(id).Error
}

// FindAll implements Repository.
func (r *repository) FindAll(story_id uint) ([]models.StoryContentRareWord, error) {
   var StoryContentRareWord []models.StoryContentRareWord

   if err := r.DB.Where("story_id = ?", story_id).Find(&StoryContentRareWord).Error; err != nil {
	   return []models.StoryContentRareWord{}, err
	}
   return StoryContentRareWord, nil
}



// Update implements Repository.
func (r *repository) Update(id uint, rareWord models.StoryContentRareWord) (*models.StoryContentRareWord, error) {
	if err := r.DB.Save(rareWord).Error; err != nil{
		return nil, err
	}
	return &rareWord, nil
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		DB: db,
	}
}
