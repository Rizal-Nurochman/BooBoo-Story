package creators

import (
	"github.com/BooBooStory/config/models"
	"gorm.io/gorm"
)

type Repository interface {
	Create(creator models.Creator) (models.Creator, error)
	FindAll(page int, limit int, q string) ([]models.Creator, int64, error)
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

func (r *repository) FindAll(page int, limit int, q string) ([]models.Creator, int64, error) {
	var creators []models.Creator
	var total int64
	query := r.db.Model(&models.Creator{})

	if q != "" {
		query = query.Where("name ILIKE ?", "%"+q+"%")
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	err = query.Offset(offset).Limit(limit).Find(&creators).Error
	if err != nil {
		return nil, 0, err
	}

	return creators, total, nil
}

func (r *repository) FindByID(creatorID uint, includes []string, page int, limit int, q string) (*models.Creator, error) {
	var creator models.Creator
	query := r.db
	for _, inc := range includes {
		query = query.Preload(inc)
	}

	err := query.First(&creator, creatorID).Error
	if err != nil {
		return nil, err
	}

	var stories []models.Story
	storyQuery := r.db.Model(&models.Story{}).Where("creator_id = ?", creatorID)

	if q != "" {
		storyQuery = storyQuery.Where("title ILIKE ?", "%"+q+"%")
	}

	offset := (page - 1) * limit
	err = storyQuery.Offset(offset).Limit(limit).Find(&stories).Error
	if err != nil {
		return nil, err
	}

	creator.Stories = stories

	return &creator, nil
}

func (r *repository) Update(creatorID uint, creator models.Creator) (*models.Creator, error) {
	var existingCreator models.Creator
	err := r.db.Where("id = ?", creatorID).First(&existingCreator).Error
	if err != nil {
		return nil, err
	}

	err = r.db.Model(&existingCreator).Updates(creator).Error
	if err != nil {
		return nil, err
	}

	return &existingCreator, nil
}

func (r *repository) Delete(creatorID uint) error {
	err := r.db.Delete(&models.Creator{}, creatorID).Error
	if err != nil {
		return err
	}

	return nil
}
