package service

import (
	"gobio/model"
)

type LinkService interface {
	AddLink(linkRequest model.AddLinkRequest, ID int) (linkResponse model.AddLinkResponse, err error)
	UpdateLink(request model.UpdateLinkRequest, ID int) (response model.UpdateLinkResponse, err error)
	DeleteLink(ID int) error
	List(username string) (allLink []model.ShowAllLinkResponse, err error)
}
