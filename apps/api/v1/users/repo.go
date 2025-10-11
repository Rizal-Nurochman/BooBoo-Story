// File: repository.go
package users

import (
	"github.com/BooBooStory/config/models"
	"gorm.io/gorm"
)

type Repository interface {
	FindByEmail(email string) (*models.User, error)
	CreateUser(user *models.User) (*models.User, error)
	FindByID(ID uint) (*models.User, error)
	FindByResetToken(token string) (*models.User, error)
	Update(user *models.User) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) CreateUser(user *models.User) (*models.User, error) {
	err :=r.db.Create(user).Error

	if err != nil{
		return nil, err
	}

	return user, nil
}

func (r *repository) FindByEmail(email string) (*models.User, error) {
   var existUser models.User
   err := r.db.Where("email = ?", email).First(&existUser).Error

   if err != nil{
	return  nil, err
   }

   return  &existUser, nil
}

func (r *repository) FindByID(ID uint) (*models.User, error) {
  var existUser models.User
  err := r.db.First(&existUser, ID).Error

  if err != nil{
		return  nil, err
  }

  return &existUser, nil
}

func (r *repository) FindByResetToken(token string) (*models.User, error) {
	var user models.User
	if token == "" {
		return nil, gorm.ErrRecordNotFound
	}
	err := r.db.Where("reset_password_token = ?", token).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *repository) Update(user *models.User) error {
	err:=r.db.Save(user).Error
	if err != nil {
		return err
	}

	return nil
}