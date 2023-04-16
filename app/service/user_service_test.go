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

var userRepository = new(mocks.UserRepository)
var userService = userServiceImpl{UserRepository: userRepository}

func TestRegister(t *testing.T) {
	userRequest := model.RegisterUserRequest{
		Name:     "name",
		Username: "username",
		Email:    "email@gmail.com",
		Password: "password",
	}

	t.Run("success", func(t *testing.T) {
		userRepository.Mock.On("Insert", mock.Anything).Return(nil).Once()
		_, err := userService.Register(userRequest)
		assert.NoError(t, err)

	})

	t.Run("failed", func(t *testing.T) {
		userRepository.Mock.On("Insert", mock.Anything).Return(errors.New("error it's happening")).Once()
		_, err := userService.Register(userRequest)
		assert.Error(t, err)

	})
}

func TestLogin(t *testing.T) {
	userRequest := model.LoginUserRequest{
		Email:    "email@email.com",
		Password: "password",
	}

	user := entity.User{
		Id:        1,
		Name:      "name",
		Username:  "username",
		Email:     "email@email.com",
		Password:  "$2a$04$ZcTZW.HCbBOdZHkhucyAzeDEwWlxJfuDwamDxS2Y4mbk3up/IgMry",
		AvatarUrl: "https://gobio.com/avatar.png",
	}

	t.Run("success", func(t *testing.T) {
		userRepository.Mock.On("FindByEmail", mock.Anything).Return(user, nil).Once()
		_, err := userService.Login(userRequest)
		assert.NoError(t, err)

	})

	t.Run("failed", func(t *testing.T) {
		userRepository.Mock.On("FindByEmail", mock.Anything).Return(user, errors.New("error it's happening")).Once()
		_, err := userService.Login(userRequest)
		assert.Error(t, err)

	})
}

func TestUploadAvatar(t *testing.T) {
	user := entity.User{
		Id:       1,
		Name:     "name",
		Username: "username",
		Email:    "email@email.com",
		Password: "$2a$04$ZcTZW.HCbBOdZHkhucyAzeDEwWlxJfuDwamDxS2Y4mbk3up/IgMry",
	}

	t.Run("success", func(t *testing.T) {
		userRepository.Mock.On("FindByID", mock.Anything).Return(user, nil).Once()
		userRepository.Mock.On("UpdateAvatar", mock.Anything).Return(user, nil).Once()
		_, err := userService.UploadAvatar(1, "https://gobio.com/avatar.png")
		assert.NoError(t, err)
	})

	t.Run("failed", func(t *testing.T) {
		userRepository.Mock.On("FindByID", mock.Anything).Return(user, errors.New("error it's happening")).Once()
		userRepository.Mock.On("UpdateAvatar", mock.Anything).Return(user, nil).Once()
		_, err := userService.UploadAvatar(1, "https://gobio.com/avatar.png")
		assert.Error(t, err)
	})
}
