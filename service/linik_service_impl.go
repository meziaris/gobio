package service

import (
	"gobio/entity"
	"gobio/model"
	"gobio/repository"
	"time"
)

type linkServiceImpl struct {
	LinkRepository repository.LinkRepository
}

func NewLinkService(repository *repository.LinkRepository) LinkService {
	return &linkServiceImpl{
		LinkRepository: *repository,
	}
}

func (service *linkServiceImpl) AddLink(linkRequest model.AddLinkRequest, ID int) (linkResponse model.AddLinkResponse, err error) {
	link := entity.Link{
		Title:     linkRequest.Title,
		Url:       linkRequest.Url,
		UserId:    ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = service.LinkRepository.Insert(link)
	if err != nil {
		return linkResponse, err
	}

	response := model.AddLinkResponse{
		Title:  link.Title,
		Url:    link.Url,
		UserId: link.UserId,
	}

	return response, nil
}
