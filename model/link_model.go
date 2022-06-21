package model

import "time"

type AddLinkRequest struct {
	Title  string `json:"title"`
	Url    string `json:"url"`
	UserId int    `json:"user_id"`
}

type AddLinkResponse struct {
	Title  string `json:"title"`
	Url    string `json:"url"`
	UserId int    `json:"user_id"`
}

type ShowAllLinkResponse struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateLinkRequest struct {
	Title string `json:"title"`
	Url   string `json:"url"`
}

type UpdateLinkResponse struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Url       string    `json:"url"`
	UserId    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
