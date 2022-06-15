package repository

import (
	"gobio/entity"

	"gorm.io/gorm"
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{
		DB: db,
	}
}

type userRepositoryImpl struct {
	DB *gorm.DB
}

func (repository *userRepositoryImpl) Insert(user entity.User) error {
	err := repository.DB.Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}
