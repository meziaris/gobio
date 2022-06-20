package service

import "gobio/model"

type LinkService interface {
	AddLink(linkRequest model.AddLinkRequest, ID int) (linkResponse model.AddLinkResponse, err error)
	List(username string) (allLink []model.ShowAllLinkResponse, err error)
}
