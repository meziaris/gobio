package entity

import "time"

type Link struct {
	Id        int       `json:"id" gorm:"primary_key;auto_increment;not_null"`
	Title     string    `json:"title" gorm:"type:varchar"`
	Url       string    `json:"url" gorm:"type:varchar"`
	UserId    int       `json:"user_id" gorm:"type:varchar"`
	User      User      `gorm:"foreignKey:UserId;references:Id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
