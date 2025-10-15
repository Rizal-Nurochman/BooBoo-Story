package bookmarks

import (
	"github.com/BooBooStory/config/models"
	"gorm.io/gorm"
)

type Repository interface {
	Create(bookmark models.StoryBookmark) (models.StoryBookmark, error)
	FindAll(userID uint,includes []string, page int, limit int, q string) ([]models.StoryBookmark, int64, error)
	FindByID(userID uint, bookmarkID uint, includes []string) (models.StoryBookmark, error)
	Delete(ID uint) (error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(bookmark models.StoryBookmark) (models.StoryBookmark, error) {
	err := r.db.Create(&bookmark).Error
	if err != nil {
		return bookmark, err
	}
	return bookmark, nil
}

func (r *repository) FindAll(userID uint, includes []string, page int, limit int, q string) ([]models.StoryBookmark, int64, error) {
	var bookmarks []models.StoryBookmark
	var total int64

	query:=r.db.Model(&models.StoryBookmark{}).Where("user_id = ?", userID)
	includes = append(includes, "Story")

	if q!=""{
		query = query.Where("title Like ?","%"+q+"%")
	}

	for _, inc := range includes{
		query =query.Preload(inc)
	}

	if err := query.Count(&total).Error; err != nil{
		return  nil, 0, err
	}

	if err :=query.Offset((page-1)*limit).Limit(limit).Find(&bookmarks).Error; err != nil{
		return nil, 0, err
	}

	return bookmarks, total, nil
}

func (r *repository) FindByID(userID uint, bookmarkID uint, includes []string) (models.StoryBookmark, error) {
	var bookmark models.StoryBookmark

	err := r.db.Where("user_id = ? AND id = ?", userID, bookmarkID).First(&bookmark).Error
	if err != nil {
		return bookmark, err
	}

	return bookmark, nil
}

func (r *repository) Delete(ID uint) (error) {
	var bookmark models.StoryBookmark
	err := r.db.Where("id = ?", ID).Delete(&bookmark).Error
	if err != nil {
		return err
	}
	return nil
}