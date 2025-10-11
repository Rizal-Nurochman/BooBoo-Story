package categories

import "github.com/BooBooStory/config/models"

type Service interface {
	Create(category models.Category) (*models.Category, error)
	GetAll(withChildren bool, page int, limit int, q string) ([]models.Category, int64, error)
	GetByID(id uint, withChildren bool) (*models.Category, error)
	Update(id uint, category models.Category) (*models.Category, error)
	Delete(id uint) error

	// 🔁 Hierarchical operations
	GetChildren(parentID uint) ([]models.Category, error)
	GetParent(childID uint) (*models.Category, error)
	CreateSubCategory(parentID uint, category models.Category) (*models.Category, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

// Create implements Service.
func (s *service) Create(category models.Category) (*models.Category, error) {
	panic("unimplemented")
}

// CreateSubCategory implements Service.
func (s *service) CreateSubCategory(parentID uint, category models.Category) (*models.Category, error) {
	panic("unimplemented")
}

// Delete implements Service.
func (s *service) Delete(id uint) error {
	panic("unimplemented")
}

// GetAll implements Service.
func (s *service) GetAll(withChildren bool, page int, limit int, q string) ([]models.Category, int64, error) {
	panic("unimplemented")
}

// GetByID implements Service.
func (s *service) GetByID(id uint, withChildren bool) (*models.Category, error) {
	panic("unimplemented")
}

// GetChildren implements Service.
func (s *service) GetChildren(parentID uint) ([]models.Category, error) {
	panic("unimplemented")
}

// GetParent implements Service.
func (s *service) GetParent(childID uint) (*models.Category, error) {
	panic("unimplemented")
}

// Update implements Service.
func (s *service) Update(id uint, category models.Category) (*models.Category, error) {
	panic("unimplemented")
}
