package bookmarks

import (
	"github.com/BooBooStory/config/models"
	"github.com/BooBooStory/utils"
)

type Service interface {
	CreateBookmark(userID uint, storyID uint) (models.StoryBookmark, error)
	GetAllBookmarks(userID uint, includes []string, page int, limit int, q string) ([]models.StoryBookmark, int64, error)
	GetBookmarkByID(userID uint, bookmarkID string, includes []string) (models.StoryBookmark, error)
	DeleteBookmark(idStr string) (error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateBookmark(userID uint, storyID uint) (models.StoryBookmark, error) {
	bookmark := models.StoryBookmark{
		UserID:  userID,
		StoryID: storyID,
	}

	newBookmark, err := s.repository.Create(bookmark)
	if err != nil {
		return newBookmark, err
	}
	return newBookmark, nil
}

func (s *service) GetAllBookmarks(userID uint, includes []string, page int, limit int, q string) ([]models.StoryBookmark, int64, error) {
	bookmarks, total, err := s.repository.FindAll(userID, includes, page, limit, q)

	if err != nil {
		return nil, 0, err
	}

	return bookmarks, total, nil
}

func (s *service) GetBookmarkByID(userID uint, idStr string, includes []string) (models.StoryBookmark, error) {
	bookmarkID, err := utils.ToUint(idStr)
	if err != nil {
		return models.StoryBookmark{}, err
	}

	bookmark, err := s.repository.FindByID(userID, bookmarkID, includes)
	if err != nil {
		return models.StoryBookmark{}, err
	}

	return bookmark, nil
}

func (s *service) DeleteBookmark(idStr string) (error) {
	id, err := utils.ToUint(idStr)
	if err != nil {
		return err
	}

	err = s.repository.Delete(uint(id))
	if err != nil {
		return err
	}

	return nil
}