package stories

import (
	"github.com/BooBooStory/config/models"
	"gorm.io/gorm"
)

type Repository interface {
	Create(story models.Story) (*models.Story, error)
	Update(id uint, story models.Story) (*models.Story, error)
	Delete(id uint) error
	FindAll(includes []string, page int, limit int, q string) ([]models.Story, int64, error)
	FindByID(id uint, includes []string) (*models.Story, error)
}

type repository struct {
	db *gorm.DB
}

func (r *repository) Create(story models.Story) (*models.Story, error) {
	if err := r.db.Create(story).Error; err != nil{
		return nil, err
	}

	return &story, nil
}

// Delete implements Repository.
func (r *repository) Delete(id uint) error {
	return r.db.Delete(&models.Story{}, id).Error
}

// FindAll implements Repository.
func (r *repository) FindAll(includes []string, page int, limit int, q string) ([]models.Story, int64, error) {
	var stories []models.Story
	var total int64

	query := r.db.Model(&models.Story{})
	includes = append(includes, "Creator")

	if q!=""{
		query = query.Where("title Like ?","%"+q+"%")
	}

	for _, inc := range includes{
		query =query.Preload(inc)
	}

	if err := query.Count(&total).Error; err != nil{
		return  nil, 0, err
	}

	if err :=query.Offset((page-1)*limit).Limit(limit).Find(&stories).Error; err != nil{
		return nil, 0, err
	}

	return stories, total, nil
}

// FindByID implements Repository.
func (r *repository) FindByID(id uint, includes []string) (*models.Story, error) {
	var story models.Story

	query :=r.db.Model(&models.Story{})
	includes = append(includes, "Creator")

	for _, inc := range includes{
		query =query.Preload(inc)
	}

	if err := query.First(&story).Error; err != nil {
		return nil, err
	}
	return &story, nil	
}

// Update implements Repository.
func (r *repository) Update(id uint, story models.Story) (*models.Story, error) {
	if err:= r.db.Save(&story).Error; err != nil{
		return  nil, err
	}
	return &story, nil
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}
