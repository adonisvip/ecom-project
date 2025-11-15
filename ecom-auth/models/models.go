package models

import "time"

type Users struct {
	ID           uint      `gorm:"primaryKey;column:id" json:"id"`
	Username     string    `gorm:"uniqueIndex;not null;size:255" json:"username"`
	Password     string    `gorm:"not null;size:255" json:"-"`
	Email        string    `gorm:"uniqueIndex;not null;size:255" json:"email"`
	Fullname     *string   `gorm:"size:255" json:"fullname,omitempty"`
	Phone        *string   `gorm:"size:20" json:"phone,omitempty"`
	Token        *string   `gorm:"size:500" json:"token,omitempty"`
	RefreshToken *string   `gorm:"column:refresh_token;size:500" json:"refresh_token,omitempty"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (Users) TableName() string { return "users" }
