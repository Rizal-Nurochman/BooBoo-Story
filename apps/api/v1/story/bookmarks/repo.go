package bookmarks

import (
	"github.com/BooBooStory/config/models"
	"gorm.io/gorm"
)

type Repository interface {
	Insert(bookmark models.StoryBookmark) (models.StoryBookmark, error)
	FindAll(userID string) ([]models.StoryBookmark, error)
	FindByID(userID string, id string) (models.StoryBookmark, error)
	Delete(bookmark models.StoryBookmark) (models.StoryBookmark, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Insert(bookmark models.StoryBookmark) (models.StoryBookmark, error) {
	err := r.db.Create(&bookmark).Error
	if err != nil {
		return bookmark, err
	}
	return bookmark, nil
}

func (r *repository) FindAll(userID string) ([]models.StoryBookmark, error) {
	var bookmarks []models.StoryBookmark
	err := r.db.Where("user_id = ?", userID).Find(&bookmarks).Error
	if err != nil {
		return bookmarks, err
	}
	return bookmarks, nil
}

func (r *repository) FindByID(userID string, id string) (models.StoryBookmark, error) {
	var bookmark models.StoryBookmark
	err := r.db.Where("user_id = ? AND id = ?", userID, id).First(&bookmark).Error
	if err != nil {
		return bookmark, err
	}
	return bookmark, nil
}

func (r *repository) Delete(bookmark models.StoryBookmark) (models.StoryBookmark, error) {
	err := r.db.Delete(&bookmark).Error
	if err != nil {
		return bookmark, err
	}
	return bookmark, nil
}