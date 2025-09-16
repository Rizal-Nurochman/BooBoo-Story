package usersrepository

import (
	"github.com/BooBooStory/config/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByEmail(email string) (*models.User, error)
	CreateUser(user *models.User) (*models.User, error)
	FindByID(ID uint) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var daftarUser models.User
	err := r.db.Where("email = ?", email).First(&daftarUser).Error
	if err != nil {
		return nil, err
	}
	return &daftarUser, nil
}

func (r *userRepository) CreateUser(user *models.User) (*models.User, error) {
	err := r.db.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) FindByID(ID uint) (*models.User, error) {
    var user models.User
    err := r.db.First(&user, ID).Error
    if err != nil {
        return nil, err
    }
    return &user, nil
}