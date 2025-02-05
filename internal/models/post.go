package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	Id        string `gorm:"primarykey"`
	UserId    string
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	Comments  []Comment
	User      User `gorm:"foreignKey:UserID"`
}

func (p *Post) BeforeCreate(db *gorm.DB) (err error) {
	p.Id = uuid.New().String()
	return
}
