package repository

import "gobio/entity"

type UserRepository interface {
	Insert(user entity.User) error
	FindByEmail(email string) (entity.User, error)
	FindByUsername(username string) (entity.User, error)
}
