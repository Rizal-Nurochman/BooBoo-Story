package creators

import (
	"github.com/BooBooStory/config/models"
	"gorm.io/gorm"
)

type Repository interface {
	Create(creator models.Creator) (models.Creator, error)
	FindAll(creatorID uint) ([]models.Creator, error)
	FindByID(creatorID uint, includes []string, page int, limit int, q string) (*models.Creator, error)
	Update(creatorID uint, creator models.Creator) (*models.Creator, error)
	Delete(creatorID uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(creator models.Creator) (models.Creator, error) {
	err := r.db.Create(&creator).Error
	if err != nil {
		return creator, err
	}

	return creator, nil
}