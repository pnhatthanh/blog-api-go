package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id        string    `gorm:"primaryKey;type:varchar(36)"`
	FullName  string    `gorm:"type:varchar(255);not null"`
	Email     string    `gorm:"unique;not null"`
	Password  string    `json:"-" gorm:"not null"`
	CreatedAt time.Time `json:"-"`

	Posts    []Post    `gorm:"foreignKey:UserId"`
	Comments []Comment `gorm:"foreignKey:UserId"`
}

func (u *User) BeforeCreate(db *gorm.DB) (err error) {
	u.Id = uuid.New().String()
	return nil
}
