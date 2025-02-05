package models

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Id        string
	UserId    string
	PostId    string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	User      User
}
