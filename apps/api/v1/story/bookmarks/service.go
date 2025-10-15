package bookmarks

import (
	"github.com/BooBooStory/config/models"
	"github.com/google/uuid"
)

type Service interface {
	CreateBookmark(userID string, storyID string) (models.StoryBookmark, error)
	GetBookmarks(userID string, page int, limit int) ([]models.StoryBookmark, int64, error)
	DeleteBookmark(userID string, id string) (models.StoryBookmark, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateBookmark(userID string, storyID string) (models.StoryBookmark, error) {
	bookmark := models.StoryBookmark{
		UserID:  uint(uuid.MustParse(userID).ID()),
		StoryID: uint(uuid.MustParse(storyID).ID()),
	}

	newBookmark, err := s.repository.Insert(bookmark)
	if err != nil {
		return newBookmark, err
	}
	return newBookmark, nil
}

func (s *service) GetBookmarks(userID string, page int, limit int) ([]models.StoryBookmark, int64, error) {
	bookmarks, total, err := s.repository.FindAll(userID, page, limit)
	if err != nil {
		return bookmarks, total, err
	}
	return bookmarks, total, nil
}

func (s *service) DeleteBookmark(userID string, id string) (models.StoryBookmark, error) {
	bookmark, err := s.repository.FindByID(userID, id)
	if err != nil {
		return bookmark, err
	}
	deletedBookmark, err := s.repository.Delete(bookmark)
	if err != nil {
		return deletedBookmark, err
	}
	return deletedBookmark, nil
}