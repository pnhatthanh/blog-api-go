package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id        string `gorm:"primaryKey"`
	FullName  string
	Email     string
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"-"`
}

func (u *User) BeforeCreate(db *gorm.DB) (err error) {
	u.Id = uuid.New().String()
	return
}
