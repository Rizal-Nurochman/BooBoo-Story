package creators

import "github.com/BooBooStory/config/models"

type UpdateCreatorInput struct {
	Name       string `json:"name" binding:"required"`
	Bio        string `json:"bio" binding:"required"`
	WebsiteURL string `json:"website_url"`
	Occupation string `json:"occupation"`
	Location   string `json:"location"`
	IgUrl      string `json:"ig_url"`
	TiktokUrl  string `json:"tiktok_url"`
	YoutubeURL string `json:"youtube_url"`
	LinkedInURL string `json:"linkedin_url"`
	Portfolio   string `json:"portfolio"`
	IsVerified  bool   `json:"is_verified"`
	Status 			string `json:"status"`
}

type Service interface {
	GetCreatorProfile(creatorID uint) (models.Creator, error)
	UpdateCreatorProfile(creatorID uint, input UpdateCreatorInput) (models.Creator, error)
	GetCreatorProfileByID(id uint) (models.Creator, error)
	DeleteCreatorByID(id uint) (models.Creator, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCreatorProfile(userID uint) (models.Creator, error) {
	creator, err := s.repository.FindByCreatorID(userID)
	if err != nil {
		return creator, err
	}
	return creator, nil
}

func (s *service) UpdateCreatorProfile(creatorID uint, input UpdateCreatorInput) (models.Creator, error) {
	creator, err := s.repository.FindByCreatorID(creatorID)
	if err != nil {
		return creator, err
	}

	creator.Name = input.Name
	creator.Bio = input.Bio
	creator.WebsiteURL = input.WebsiteURL
	creator.Occupation = input.Occupation
	creator.Location = input.Location
	creator.IgUrl = input.IgUrl
	creator.TiktokUrl = input.TiktokUrl
	creator.YoutubeURL = input.YoutubeURL
	creator.LinkedInURL = input.LinkedInURL
	creator.Portfolio = input.Portfolio

	updatedCreator, err := s.repository.Update(creator)
	if err != nil {
		return updatedCreator, err
	}
	return updatedCreator, nil
}

func (s *service) GetCreatorProfileByID(id uint) (models.Creator, error) {
	creator, err := s.repository.FindByID(id)
	if err != nil {
		return creator, err
	}
	return creator, nil
}

func (s *service) DeleteCreatorByID(id uint) (models.Creator, error) {
	creator, err := s.repository.FindByID(id)
	if err != nil {
		return creator, err
	}

	deletedCreator, err := s.repository.Delete(creator)
	if err != nil {
		return deletedCreator, err
	}
	return deletedCreator, nil
}