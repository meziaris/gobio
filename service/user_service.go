package service

import "gobio/model"

type UserService interface {
	Register(request model.RegisterUserRequest) (model.RegisterUserResponse, error)
	Login(request model.LoginUserRequest, token string) (model.LoginUserResponse, error)
}
