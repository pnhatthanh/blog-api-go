package repositories

import (
	"blogapi/internal/dtos"
	"blogapi/internal/models"

	"gorm.io/gorm"
)

type PostRepository interface {
	GetAll(offset, limit int) []*models.Post
	GetById(id string) (*models.Post, error)
	Delete(id string) error
	Create(post *dtos.CreatePostInput) *models.Post
}
type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *postRepository {
	return &postRepository{db: db}
}

func (repo *postRepository) GetAll(offset, limit int) []*models.Post {
	var postResponse []*models.Post
	repo.db.Preload("User").Preload("Comments", func(db *gorm.DB) *gorm.DB {
		return db.Limit(5).Order("created_at desc")
	}).Find(&postResponse)
	return postResponse
}
func (repo *postRepository) GetById(id string) (*models.Post, error) {
	var postResponse models.Post
	err := repo.db.Preload("User").Preload("Comments", func(db *gorm.DB) *gorm.DB {
		return db.Limit(5).Order("created_at desc")
	}).First(&postResponse, "id=?", id).Error
	if err != nil {
		return nil, err
	}
	return &postResponse, nil
}
func (repo *postRepository) Delete(id string) error {
	var post models.Post
	if err := repo.db.First(&post, "id = ?", id).Error; err != nil {
		return err
	}
	if err := repo.db.Delete(&post).Error; err != nil {
		return err
	}
	return nil
}
func (repo *postRepository) Create(post *dtos.CreatePostInput) *models.Post {
	_post := models.Post{
		Title:   post.Title,
		Content: post.Content,
		UserId:  post.UserId,
	}
	repo.db.Create(&_post)
	return &_post
}
