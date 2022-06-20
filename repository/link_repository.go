package repository

import "gobio/entity"

type LinkRepository interface {
	Insert(link entity.Link) error
	Delete(link entity.Link) error
	FindAllUserLink(id int) (link []entity.Link, err error)
}
