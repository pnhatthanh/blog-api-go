package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	Id        string    `gorm:"primaryKey;type:varchar(36)"`
	UserId    string    `gorm:"not null"`
	Title     string    `gorm:"type:varchar(255);not null"`
	Content   string    `gorm:"type:text;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time

	User     User      `gorm:"foreignKey:UserId"`
	Comments []Comment `gorm:"foreignKey:PostId"`
}

func (p *Post) BeforeCreate(db *gorm.DB) (err error) {
	p.Id = uuid.New().String()
	return nil
}
