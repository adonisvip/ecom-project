package models

import "time"

type Users struct {
	UserID        *int `json:"user_id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Fullname      *string
	Username      *string
	Password      *string
	Email         *string
	Phone         *string
	Token         *string
	Refresh_Token *string
	Created_At    time.Time
	Updated_At    time.Time
}

func (Users) TableName() string { return "users" }