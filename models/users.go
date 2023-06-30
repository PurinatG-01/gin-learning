package model

import (
	"time"
)

type Users struct {
	Id            int        `gorm:"id" json:"id"`
	Username      string     `gorm:"username" json:"username" form:"username"`
	DisplayName   string     `gorm:"display_name" json:"displayName" form:"displayName"`
	DisplayImgUrl string     `gorm:"display_img_url" json:"displayImgUrl" form:"displayImgUrl"`
	IsAdmin       bool       `gorm:"is_admin" json:"isAdmin"`
	Email         string     `gorm:"email" json:"email" form:"email"`
	CreatedAt     *time.Time `gorm:"created_at" json:"createdAt"`
	UpdatedAt     *time.Time `gorm:"updated_at" json:"updatedAt"`
	Password      string     `gorm:"password" json:"password"`
	TotalMoney    int        `gorm:"total_money" json:"totalMoney"`
}

type PublicUser struct {
	Id            int        `json:"id"`
	Username      string     `json:"username"`
	DisplayName   string     `json:"displayName"`
	DisplayImgUrl string     `json:"displayImgUrl"`
	Email         string     `json:"email"`
	CreatedAt     *time.Time `json:"createdAt"`
	UpdatedAt     *time.Time `json:"updatedAt"`
}

type FormUser struct {
	Username      string `json:"username" form:"username"`
	DisplayName   string `json:"displayName" form:"displayName"`
	DisplayImgUrl string `json:"displayImgUrl" form:"displayImgUrl"`
	Email         string `json:"email" form:"email"`
	Password      string `json:"password" form:"password"`
}
