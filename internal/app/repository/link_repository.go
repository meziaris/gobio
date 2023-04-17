package repository

import "gobio/internal/app/entity"

type LinkRepository interface {
	Insert(link entity.Link) error
	FindLinkById(ID int) (entity.Link, error)
	Update(link entity.Link) (entity.Link, error)
	DeleteLinkById(link entity.Link) error
	FindAllUserLink(id int) (link []entity.Link, err error)
}
