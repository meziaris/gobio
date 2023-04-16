package repository

import (
	"gobio/app/entity"

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

func (repository *linkRepositoryImpl) FindLinkById(ID int) (entity.Link, error) {
	var link entity.Link
	err := repository.DB.Where("id = ?", ID).First(&link).Error
	if err != nil {
		return link, err
	}

	return link, nil
}

func (repository *linkRepositoryImpl) Update(link entity.Link) (entity.Link, error) {
	err := repository.DB.Save(&link).Error
	if err != nil {
		return link, err
	}

	return link, nil
}

func (repository *linkRepositoryImpl) DeleteLinkById(link entity.Link) error {
	err := repository.DB.Delete(&link).Error
	if err != nil {
		return err
	}

	return nil
}

func (repository *linkRepositoryImpl) FindAllUserLink(id int) (links []entity.Link, err error) {
	var newLinks []entity.Link
	err = repository.DB.Joins("left join users ON users.id = links.id").Preload("User").Where("links.user_id = ?", id).
		Find(&newLinks).Error
	if err != nil {
		return links, err
	}

	for _, link := range newLinks {
		links = append(links, entity.Link{
			Id:        link.Id,
			Title:     link.Title,
			Url:       link.Url,
			CreatedAt: link.CreatedAt,
			UpdatedAt: link.UpdatedAt,
		})
	}

	return links, nil
}
