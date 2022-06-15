package repository

import "gobio/entity"

type UserRepository interface {
	Insert(user entity.User) error
}
