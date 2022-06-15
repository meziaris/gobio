package entity

import "time"

type User struct {
	Id        int       `json:"id" gorm:"primary_key;auto_increment;not_null"`
	Name      string    `json:"name" gorm:"type:varchar"`
	Username  string    `json:"username" gorm:"unique;type:varchar"`
	Email     string    `json:"email" gorm:"type:varchar"`
	Password  string    `json:"password" gorm:"type:varchar"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
