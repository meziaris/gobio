package repository

import "gobio/internal/app/entity"

type UserRepository interface {
	Insert(user entity.User) error
	FindByID(id int) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
	FindByUsername(username string) (entity.User, error)
	UpdateAvatar(user entity.User) (entity.User, error)
}
