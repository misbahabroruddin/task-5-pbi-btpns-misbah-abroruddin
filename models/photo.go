package models

import (
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	Title string `json:"title"`
	Caption string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	UserID string `json:"user_id"`
	User User `gorm:"foreignkey:UserID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
}

