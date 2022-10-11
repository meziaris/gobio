package service

import (
	"errors"
	"gobio/app/entity"
	"gobio/app/model"
	"gobio/app/repository/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepositoryLink = new(mocks.UserRepository)
var linkRepository = new(mocks.LinkRepository)
var linkService = linkServiceImpl{LinkRepository: linkRepository, UserRepository: userRepositoryLink}

func TestAddLink(t *testing.T) {
	link := model.AddLinkRequest{
		Title:  "link",
		Url:    "link.com",
		UserId: 1,
	}

	t.Run("success", func(t *testing.T) {
		linkRepository.Mock.On("Insert", mock.Anything).Return(nil).Once()

		_, err := linkService.AddLink(link, 1)
		assert.NoError(t, err)
	})

	t.Run("failed", func(t *testing.T) {
		linkRepository.Mock.On("Insert", mock.Anything).Return(errors.New("error it's happening")).Once()

		_, err := linkService.AddLink(link, 1)
		assert.Error(t, err)
	})
}

func TestUpdateLink(t *testing.T) {
	link := entity.Link{
		Id:     1,
		Title:  "title",
		Url:    "url.com",
		UserId: 1,
	}
	linkRequest := model.UpdateLinkRequest{
		Title: "title",
		Url:   "url.com",
	}

	t.Run("success", func(t *testing.T) {
		linkRepository.Mock.On("FindLinkById", mock.Anything).Return(link, nil).Once()
		linkRepository.Mock.On("Update", mock.Anything).Return(link, nil).Once()

		_, err := linkService.UpdateLink(linkRequest, 1, 1)
		assert.NoError(t, err)
	})

	t.Run("failed", func(t *testing.T) {
		linkRepository.Mock.On("FindLinkById", mock.Anything).Return(link, errors.New("error it's happening")).Once()
		linkRepository.Mock.On("Update", mock.Anything).Return(link, errors.New("error it's happening")).Once()

		_, err := linkService.UpdateLink(linkRequest, 1, 1)
		assert.Error(t, err)
	})
}

func TestDeleteLink(t *testing.T) {
	link := entity.Link{
		Id:     1,
		Title:  "title",
		Url:    "url.com",
		UserId: 1,
	}

	t.Run("success", func(t *testing.T) {
		linkRepository.Mock.On("FindLinkById", mock.Anything).Return(link, nil).Once()
		linkRepository.Mock.On("DeleteLinkById", mock.Anything).Return(nil).Once()

		err := linkService.DeleteLink(1, 1)
		assert.NoError(t, err)
	})

	t.Run("failed", func(t *testing.T) {
		linkRepository.Mock.On("FindLinkById", mock.Anything).Return(link, errors.New("error it's happening")).Once()
		linkRepository.Mock.On("DeleteLinkById", mock.Anything).Return(errors.New("error it's happening")).Once()

		err := linkService.DeleteLink(1, 1)
		assert.Error(t, err)
	})
}

func TestListAllLink(t *testing.T) {
	user := entity.User{
		Id:       1,
		Name:     "name",
		Username: "username",
		Email:    "email@email.com",
	}
	links := []entity.Link{}

	t.Run("success", func(t *testing.T) {
		userRepositoryLink.Mock.On("FindByUsername", mock.Anything).Return(user, nil).Once()
		linkRepository.Mock.On("FindAllUserLink", mock.Anything).Return(links, nil).Once()

		_, err := linkService.List("username")
		assert.NoError(t, err)
	})

	t.Run("success", func(t *testing.T) {
		userRepositoryLink.Mock.On("FindByUsername", mock.Anything).Return(user, errors.New("error it's happening")).Once()
		linkRepository.Mock.On("FindAllUserLink", mock.Anything).Return(links, errors.New("error it's happening")).Once()

		_, err := linkService.List("username")
		assert.Error(t, err)
	})
}
