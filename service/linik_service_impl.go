package service

import (
	"errors"
	"gobio/entity"
	"gobio/model"
	"gobio/repository"
	"time"
)

type linkServiceImpl struct {
	LinkRepository repository.LinkRepository
	UserRepository repository.UserRepository
}

func NewLinkService(linkRepository *repository.LinkRepository, userRepository *repository.UserRepository) LinkService {
	return &linkServiceImpl{
		LinkRepository: *linkRepository,
		UserRepository: *userRepository,
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

func (service *linkServiceImpl) UpdateLink(request model.UpdateLinkRequest, ID int) (response model.UpdateLinkResponse, err error) {
	var link = entity.Link{}
	link, err = service.LinkRepository.FindLinkById(ID)

	link.Title = request.Title
	link.Url = request.Url
	link.UpdatedAt = time.Now()

	if err != nil {
		return response, nil
	}

	newLink, err := service.LinkRepository.Update(link)
	if err != nil {
		return response, nil
	}

	linkResponse := model.UpdateLinkResponse{
		Id:        newLink.Id,
		Title:     newLink.Title,
		Url:       newLink.Url,
		CreatedAt: newLink.CreatedAt,
		UpdatedAt: newLink.UpdatedAt,
	}

	return linkResponse, nil

}
func (service *linkServiceImpl) DeleteLink(ID int) error {
	var link = entity.Link{}
	link, err := service.LinkRepository.FindLinkById(ID)
	if err != nil {
		return err
	}

	err = service.LinkRepository.DeleteLinkById(link)
	if err != nil {
		return err
	}

	return nil
}

func (service *linkServiceImpl) List(username string) (response []model.ShowAllLinkResponse, err error) {
	user, err := service.UserRepository.FindByUsername(username)

	if err != nil {
		return response, err
	}

	if user.Id == 0 {
		return response, errors.New("user doesn't exist")
	}

	links, err := service.LinkRepository.FindAllUserLink(user.Id)
	if err != nil {
		return response, err
	}

	for _, link := range links {
		response = append(response, model.ShowAllLinkResponse{
			Id:        link.Id,
			Title:     link.Title,
			Url:       link.Url,
			CreatedAt: link.CreatedAt,
			UpdatedAt: link.UpdatedAt,
		})
	}

	return response, nil

}
