package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Id        string
	UserId    string
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	Comments  []Comment
}
