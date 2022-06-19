package model

type AddLinkRequest struct {
	Title  string `json:"title" gorm:"type:varchar"`
	Url    string `json:"url" gorm:"type:varchar"`
	UserId int    `json:"user_id"`
}

type AddLinkResponse struct {
	Title  string `json:"title" gorm:"type:varchar"`
	Url    string `json:"url" gorm:"type:varchar"`
	UserId int    `json:"user_id"`
}
