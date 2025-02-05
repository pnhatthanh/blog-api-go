package repositories

import (
	"blogapi/internal/dtos"
	"blogapi/internal/models"

	"gorm.io/gorm"
)

type PostRepository interface {
	GetAll(offset, limit string) []*models.Post
	GetById(id string) (*models.Post, error)
	Delete(id string) error
	Create(post *dtos.CreatePostInput) *models.Post
}
type postRepository struct{
	db *gorm.DB
}
func NewPostRepository(db *gorm.DB) *postRepository{
	return &postRepository{db: db}
}


