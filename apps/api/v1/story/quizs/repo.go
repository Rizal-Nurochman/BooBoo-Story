package quizs

import (
	"github.com/BooBooStory/config/models"
	"gorm.io/gorm"
)

type Repository interface {
	Create(quiz models.Quiz) (*models.Quiz, error)
	Delete(id uint) error
}

type repository struct {
	DB *gorm.DB
}

// Create implements Repository.
func (r *repository) Create(quiz models.Quiz) (*models.Quiz, error) {
	 if err := r.DB.Create(quiz).Error; err != nil{
		 return nil, err
	 }
	 return  &quiz, nil
}

// Delete implements Repository.
func (r *repository) Delete(id uint) error {
	return r.DB.Delete(id).Error
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		DB: db,
	}
}
