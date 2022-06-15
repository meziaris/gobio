package service

import (
	"errors"
	"gobio/entity"
	"gobio/model"
	"gobio/repository"

	"golang.org/x/crypto/bcrypt"
)

func NewUserService(repository *repository.UserRepository) UserService {
	return &userServiceImpl{
		UserRepository: *repository,
	}
}

type userServiceImpl struct {
	UserRepository repository.UserRepository
}

func (service *userServiceImpl) Register(request model.RegisterUserRequest) (response model.RegisterUserResponse, err error) {
	user := entity.User{}
	user.Name = request.Name
	user.Username = request.Username
	user.Email = request.Email
	if request.Password == "" {
		return response, errors.New("Password can't empty")
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
		Token:    "tokentoken",
	}

	return userResponse, nil

}
