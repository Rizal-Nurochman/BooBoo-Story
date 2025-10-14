package rarewords

import (
	"errors"

	"github.com/BooBooStory/config/models"
	"github.com/BooBooStory/config/validations"
	"github.com/BooBooStory/utils"
)

type Service interface {
	Create(rareWordCreate validations.StoryContentRareWordCreate) (*models.StoryContentRareWord, error)
	Update(idStr string, rareWordUpdate validations.StoryContentRareWordUpdate) (*models.StoryContentRareWord, error)
	Delete(idStr string) error
	GetAll(storyContentIDStr string) ([]models.StoryContentRareWord, error)
}

type service struct {
	repo Repository
}

func (s *service) Create(rareWordCreate validations.StoryContentRareWordCreate) (*models.StoryContentRareWord, error) {
	newRareWord := models.StoryContentRareWord{
		Title:          rareWordCreate.Title,
		Description:    rareWordCreate.Description,
		StoryContentID: rareWordCreate.StoryContentID,
	}
	createdRareWord, err := s.repo.Create(newRareWord)
	if err != nil {
		return nil, err
	}
	return createdRareWord, nil
}

func (s *service) Update(idStr string, rareWordUpdate validations.StoryContentRareWordUpdate) (*models.StoryContentRareWord, error) {
	id, err := utils.ToUint(idStr)
	if err != nil {
		return nil, err
	}

	existing, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return nil, errors.New("rare word not found")
	}

	updatedData := models.StoryContentRareWord{}

	if rareWordUpdate.Title != nil {
		updatedData.Title = *rareWordUpdate.Title
	}
	if rareWordUpdate.Description != nil {
		updatedData.Description = *rareWordUpdate.Description
	}

	return s.repo.Update(id, updatedData)
}

func (s *service) Delete(idStr string) error {
	id, err := utils.ToUint(idStr)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

func (s *service) GetAll(storyContentIDStr string) ([]models.StoryContentRareWord, error) {
	storyContentID, err := utils.ToUint(storyContentIDStr)
	if err != nil {
		return nil, err
	}
	return s.repo.FindAll(storyContentID)
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}
