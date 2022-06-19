package repository

import (
	"gobio/entity"

	"gorm.io/gorm"
)

type linkRepositoryImpl struct {
	DB *gorm.DB
}

func NewLinkRepository(db *gorm.DB) LinkRepository {
	return &linkRepositoryImpl{
		DB: db,
	}
}

func (repository *linkRepositoryImpl) Insert(link entity.Link) error {
	err := repository.DB.Create(&link).Error
	if err != nil {
		return err
	}

	return nil
}

func (repository *linkRepositoryImpl) Delete(link entity.Link) error {
	err := repository.DB.Delete(&link).Error
	if err != nil {
		return err
	}

	return nil
}
