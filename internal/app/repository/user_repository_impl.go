package repository

import (
	"gobio/internal/app/entity"

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

func (repository *userRepositoryImpl) FindByID(ID int) (entity.User, error) {
	var user entity.User
	err := repository.DB.Where("id = ?", ID).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (repository *userRepositoryImpl) FindByEmail(email string) (entity.User, error) {
	var user entity.User
	err := repository.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (repository *userRepositoryImpl) FindByUsername(username string) (entity.User, error) {
	var user entity.User
	err := repository.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (repository *userRepositoryImpl) UpdateAvatar(user entity.User) (entity.User, error) {
	err := repository.DB.Save(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
