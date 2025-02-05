package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id        string
	FullName  string
	Email     string
	Password  string
	CreatedAt time.Time
	Posts     []Post
}
