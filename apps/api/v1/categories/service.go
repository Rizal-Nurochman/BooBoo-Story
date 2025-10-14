package categories

import (
	"errors"

	"github.com/BooBooStory/config/models"
	"github.com/BooBooStory/config/validations"
)

type Service interface {
	Create(input validations.CategoryCreate) (*models.Category, error)
	GetAll(withChildren bool, page int, limit int, q string) ([]models.Category, int64, error)
	GetByID(id uint, withChildren bool) (*models.Category, error)
	Update(id uint, input validations.CategoryUpdate) (*models.Category, error)
	Delete(id uint) error

	GetChildren(parentID uint) ([]models.Category, error)
	GetParent(childID uint) (*models.Category, error)
	CreateSubCategory(parentID uint, input validations.SubCategoryCreate) (*models.Category, error)
}

type service struct {
	repo Repository
}

// Create implements Service.
func (s *service) Create(input validations.CategoryCreate) (*models.Category, error) {
	category := models.Category{
		Name: input.Name,
		ParentID: input.ParentID,
	}
	newCategory , err := s.repo.Create(category)
	 if err != nil {
		return  nil, err
	 }

	 return  newCategory, nil
}

// CreateSubCategory implements Service.
func (s *service) CreateSubCategory(parentID uint, input validations.SubCategoryCreate) (*models.Category, error) {
	if input.ParentID ==nil{
				return nil, errors.New("parent_id is required for subcategory")
	}

	category := models.Category{
		Name: input.Name,
		ParentID: &parentID,
	}

	newCategory, err := s.repo.CreateSubCategory(parentID, category)

	if err != nil{
		return nil, err
	}

	return newCategory, nil

}

// Delete implements Service.
func (s *service) Delete(id uint) error {
	return  s.repo.Delete(id)
}

// GetAll implements Service.
func (s *service) GetAll(withChildren bool, page int, limit int, q string) ([]models.Category, int64, error) {
	return s.repo.FindAll(withChildren, page, limit, q)
}

// GetByID implements Service.
func (s *service) GetByID(id uint, withChildren bool) (*models.Category, error) {
	category, err := s.repo.FindByID(id, withChildren)

	if err != nil{
		return nil, err
	}

	return category, nil
}

// GetChildren implements Service.
func (s *service) GetChildren(parentID uint) ([]models.Category, error) {
	return s.repo.FindChildren(parentID)
}

// GetParent implements Service.
func (s *service) GetParent(childID uint) (*models.Category, error) {
	category, err := s.repo.FindParent(childID)

	if err != nil{
		return  nil, err
	}

	return category, nil
}

// Update implements Service.
func (s *service) Update(id uint, input validations.CategoryUpdate) (*models.Category, error) {
	exist, err := s.repo.FindByID(id, false)
	
	if err != nil{
		return nil, err
	}

	if input.Name != ""{
		exist.Name=input.Name
	}

	if input.ParentID != nil{
		exist.ParentID = input.ParentID
	}

	updated, err := s.repo.Update(*exist)

	if err != nil{
		return  nil, err
	}

	return updated, nil
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}
