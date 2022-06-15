package service

import "gobio/model"

type UserService interface {
	Register(request model.RegisterUserRequest) (model.RegisterUserResponse, error)
}
