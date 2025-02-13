package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comment struct {
	Id        string    `gorm:"primaryKey;type:varchar(36)"`
	UserId    string    `gorm:"not null"`
	PostId    string    `gorm:"not null"`
	Content   string    `gorm:"type:text;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time

	User User `gorm:"foreignKey:UserId"`
	Post Post `gorm:"foreignKey:PostId"`
}

func (c *Comment) BeforeCreate(db *gorm.DB) (err error) {
	c.Id = uuid.New().String()
	return nil
}
