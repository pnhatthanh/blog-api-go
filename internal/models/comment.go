package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comment struct {
	Id        string `gorm:"primaryKey"`
	UserId    string
	PostId    string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	User      User `gorm:"foreignKey"`
}

func (c *Comment) BeforeCreate(db *gorm.DB) (err error) {
	c.Id = uuid.New().String()
	return
}
