package service

import (
	"errors"
	"gobio/app/entity"
	"gobio/app/model"
	"gobio/app/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(request model.RegisterUserRequest) (model.RegisterUserResponse, error)
	Login(request model.LoginUserRequest) (model.LoginUserResponse, error)
	UploadAvatar(userID int, avatarURL string) (model.UpdateAvatarResponse, error)
}

type userServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserService(repository *repository.UserRepository) UserService {
	return &userServiceImpl{
		UserRepository: *repository,
	}
}

func (service *userServiceImpl) Register(request model.RegisterUserRequest) (response model.RegisterUserResponse, err error) {
	user := entity.User{}
	user.Name = request.Name
	user.Username = request.Username
	user.Email = request.Email
	if request.Password == "" {
		return response, errors.New("password can't empty")
	}
	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)
	if err != nil {
		return response, err
	}

	user.Password = string(password)

	err = service.UserRepository.Insert(user)
	if err != nil {
		return response, err
	}

	userResponse := model.RegisterUserResponse{
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
	}

	return userResponse, nil
}

func (service *userServiceImpl) Login(request model.LoginUserRequest) (response model.LoginUserResponse, err error) {
	email := request.Email
	password := request.Password

	user, err := service.UserRepository.FindByEmail(email)
	if err != nil {
		return response, err
	}

	if user.Id == 0 {
		return response, errors.New("user doesn't exist")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return response, errors.New("incorrect password")
	}

	userResponse := model.LoginUserResponse{
		ID:       user.Id,
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
	}

	return userResponse, nil

}

func (service *userServiceImpl) UploadAvatar(userID int, avatarURL string) (response model.UpdateAvatarResponse, err error) {
	var user = entity.User{}
	user, err = service.UserRepository.FindByID(userID)
	if err != nil {
		return response, err
	}

	user.AvatarUrl = avatarURL

	newUser, err := service.UserRepository.UpdateAvatar(user)
	if err != nil {
		return response, err
	}

	response = model.UpdateAvatarResponse{
		ID:        newUser.Id,
		Name:      newUser.Name,
		Username:  newUser.Username,
		Email:     newUser.Email,
		AvatarUrl: newUser.AvatarUrl,
	}

	return response, nil
}
