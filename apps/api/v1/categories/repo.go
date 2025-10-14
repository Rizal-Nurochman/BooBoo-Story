package categories

import (
	"github.com/BooBooStory/config/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	Create(category models.Category) (*models.Category, error)
	FindAll(withChildren bool, page int, limit int, q string) ([]models.Category, int64, error)
	FindByID(id uint, withChildren bool) (*models.Category, error)
	Update(category models.Category) (*models.Category, error)
	Delete(id uint) error

	FindChildren(parentID uint) ([]models.Category, error)
	FindParent(childID uint) (*models.Category, error)
	CreateSubCategory(parentID uint, category models.Category) (*models.Category, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

// 🔹 Create
func (r *repository) Create(category models.Category) (*models.Category, error) {
	if err := r.db.Create(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

// 🔹 CreateSubCategory
func (r *repository) CreateSubCategory(parentID uint, category models.Category) (*models.Category, error) {
	category.ParentID = &parentID

	err := r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "name"}, {Name: "parent_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"name"}), // updated_at otomatis diisi GORM
	}).Create(&category).Error

	if err != nil {
		return nil, err
	}

	return &category, nil
}

// 🔹 Delete
func (r *repository) Delete(id uint) error {
	return r.db.Delete(&models.Category{}, id).Error
}

   
func (r *repository) FindAll(withChildren bool, page, limit int, q string) ([]models.Category, int64, error) {
	var (
		categories []models.Category
		total      int64
	)

	query := r.db.Model(&models.Category{})

	if q != "" {
		query = query.Where("name ILIKE ?", "%"+q+"%")
	}

	query.Count(&total)

	if withChildren {
		query = query.Preload("Children.Children")
	}

	if page > 0 && limit > 0 {
		offset := (page - 1) * limit
		query = query.Offset(offset).Limit(limit)
	}

	err := query.Find(&categories).Error
	if err != nil {
		return nil, 0, err
	}

	return categories, total, nil
}

func (r *repository) FindByID(id uint, withChildren bool) (*models.Category, error) {
	var category models.Category

	query := r.db.Model(&models.Category{})

	if withChildren {
		query = query.Preload("Children.Children")
	}

	err := query.First(&category, id).Error
	if err != nil {
		return nil, err
	}

	return &category, nil
}

// 🔹 FindChildren
func (r *repository) FindChildren(parentID uint) ([]models.Category, error) {
	var categories []models.Category
	err := r.db.Where("parent_id = ?", parentID).Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

// 🔹 FindParent
func (r *repository) FindParent(childID uint) (*models.Category, error) {
	var child models.Category
	err := r.db.First(&child, childID).Error
	if err != nil {
		return nil, err
	}

	if child.ParentID == nil {
		return nil, nil 
	}

	var parent models.Category
	err = r.db.First(&parent, *child.ParentID).Error
	if err != nil {
		return nil, err
	}

	return &parent, nil
}

// 🔹 Update
func (r *repository) Update(category models.Category) (*models.Category, error) {
	err := r.db.Model(&models.Category{}).Where("id = ?", category.ID).Updates(map[string]interface{}{
		"name":      category.Name,
		"parent_id": category.ParentID,
	}).Error

	if err != nil {
		return nil, err
	}

	return &category, nil
}
