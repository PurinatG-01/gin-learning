package model

import "time"

type User struct {
	Id              int        `gorm:"id" json:"id"`
	Username        string     `gorm:"username" json:"username"`
	DisplayName     string     `gorm:"display_name" json:"displayName"`
	DisplayImageUrl string     `gorm:"display_image_url" json:"displayImageUrl"`
	IsAdmin         bool       `gorm:"is_admin" json:"isAdmin"`
	Email           string     `gorm:"email" json:"email"`
	CreatedAt       *time.Time `gorm:"created_at" json:"createdAt"`
	UpdatedAt       *time.Time `gorm:"updated_at" json:"updatedAt"`
}
