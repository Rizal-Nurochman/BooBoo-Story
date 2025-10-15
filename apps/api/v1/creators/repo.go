package creators

import (
	"github.com/BooBooStory/config/models"
	"gorm.io/gorm"
)

type Repository interface {
	FindByCreatorID(creatorID uint) (models.Creator, error)
	Update(creator models.Creator) (models.Creator, error)
	FindByID(id uint) (models.Creator, error)
	Delete(creator models.Creator) (models.Creator, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindByCreatorID(creatorID uint) (models.Creator, error) {
	var creator models.Creator
	err := r.db.Where("user_id = ?", creatorID).First(&creator).Error
	if err != nil {
		return creator, err
	}
	return creator, nil
}

func (r *repository) Update(creator models.Creator) (models.Creator, error) {
	err := r.db.Save(&creator).Error
	if err != nil {
		return creator, err
	}
	return creator, nil
}

func (r *repository) FindByID(id uint) (models.Creator, error) {
	var creator models.Creator
	err := r.db.First(&creator, id).Error
	if err != nil {
		return creator, err
	}
	return creator, nil
}

func (r *repository) Delete(creator models.Creator) (models.Creator, error) {
	err := r.db.Delete(&creator).Error
	if err != nil {
		return creator, err
	}
	return creator, nil
}