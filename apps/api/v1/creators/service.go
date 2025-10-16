package creators

import (
	"github.com/BooBooStory/config/models"
	"github.com/BooBooStory/config/validations"
)

type service struct {
	repository Repository
}

type Service interface {
	Create(creator validations.CreatorCreate) (models.Creator, error)
	GetAll(page int, limit int, q string) ([]models.Creator, int64, error)
	GetByID(creatorID uint, includes []string, page int, limit int, q string) (*models.Creator, error)
	Update(creatorID uint, creator validations.CreatorUpdate) (*models.Creator, error)
	Delete(creatorID uint) error
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Create(creator validations.CreatorCreate) (models.Creator, error) {
	newCreator := models.Creator{
		Name:        		creator.Name,
		Bio:       			creator.Bio,
		ProfileImage:  	creator.ProfileImage,
		BannerImage:  	creator.BannerImage,
		Occupation:   	creator.Occupation,
		Location:     	creator.Location,
		WebsiteURL:   	creator.WebsiteURL,
		IgUrl:       		creator.IgUrl,
		TiktokUrl:   		creator.TiktokUrl,
		YoutubeURL:  		creator.YoutubeURL,
		LinkedInURL: 		creator.LinkedInURL,
		Portfolio:    	creator.Portfolio,
	}

	createdCreator, err := s.repository.Create(newCreator)
	if err != nil {
		return createdCreator, err
	}

	return createdCreator, nil
}

func (s *service) GetAll(page int, limit int, q string) ([]models.Creator, int64, error) {
	creators, total, err := s.repository.FindAll(page, limit, q)
	if err != nil {
		return nil, 0, err
	}

	return creators, total, nil
}

func (s *service) GetByID(creatorID uint, includes []string, page int, limit int, q string) (*models.Creator, error) {
	creator, err := s.repository.FindByID(creatorID, includes, page, limit, q)
	if err != nil {
		return nil, err
	}

	return creator, nil
}

func (s *service) Update(creatorID uint, creator validations.CreatorUpdate) (*models.Creator, error) {
	updatedCreator := models.Creator{
		Name:         	creator.Name,
		Bio: 						creator.Bio,
		ProfileImage:  	creator.ProfileImage,
		BannerImage:  	creator.BannerImage,
		Occupation:   	creator.Occupation,
		Location:     	creator.Location,
		WebsiteURL:   	creator.WebsiteURL,
		IgUrl:       		creator.IgUrl,
		TiktokUrl:   		creator.TiktokUrl,
		YoutubeURL:  		creator.YoutubeURL,
		LinkedInURL: 		creator.LinkedInURL,
		Portfolio:    	creator.Portfolio,
	}

	creatorUpdated, err := s.repository.Update(creatorID, updatedCreator)
	if err != nil {
		return nil, err
	}

	return creatorUpdated, nil
}

func (s *service) Delete(creatorID uint) error {
	err := s.repository.Delete(creatorID)
	if err != nil {
		return err
	}

	return nil
}