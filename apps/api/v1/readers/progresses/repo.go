package progresses

import (
	"github.com/BooBooStory/config/models"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
	Create(progress models.ProgressRead) (models.ProgressRead, error)
	FindAll(userID uint, includes []string, page int, limit int, q string) ([]models.ProgressRead, int64, error)
	FindByID(progressID uint, includes []string) (*models.ProgressRead, error)
	Update(progressID uint, progress models.ProgressRead) (*models.ProgressRead, error)
	Delete(progressID uint) error
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(progress models.ProgressRead) (models.ProgressRead, error) {
	err := r.db.Create(&progress).Error
	if err != nil {
		return progress, err
	}

	return progress, nil
}

func (r *repository) FindAll(userID uint, includes []string, page int, limit int, q string) ([]models.ProgressRead, int64, error) {
	var progresses []models.ProgressRead
	var total int64
	query := r.db.Model(&models.ProgressRead{}).Where("user_id = ?", userID)
	includes = append(includes, "Story")

	if q != "" {
		query = query.Joins("JOIN stories ON stories.id = progress_reads.story_id").
			Where("stories.title ILIKE ?", "%"+q+"%")
	}

	for _, inc := range includes {
		query = query.Preload(inc)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	err = query.Offset(offset).Limit(limit).Find(&progresses).Error

	if err != nil {
		return nil, 0, err
	}

	return progresses, total, nil
}

func (r *repository) FindByID(progressID uint, includes []string) (*models.ProgressRead, error) {
	var progress models.ProgressRead
	query := r.db
	for _, inc := range includes {
		query = query.Preload(inc)
	}

	err := query.First(&progress, progressID).Error
	if err != nil {
		return nil, err
	}

	return &progress, nil
}

func (r *repository) Update(progressID uint, progress models.ProgressRead) (*models.ProgressRead, error) {
	err := r.db.Model(&models.ProgressRead{}).Where("id = ?", progressID).Updates(progress).Error
	if err != nil {
		return nil, err
	}

	return &progress, nil
}

func (r *repository) Delete(progressID uint) error {
	return r.db.Delete(&models.ProgressRead{}, progressID).Error
}