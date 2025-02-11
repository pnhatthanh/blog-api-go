package services

import (
	"blogapi/internal/dtos"
	"blogapi/internal/models"
	"blogapi/internal/repositories"
)

type PostService interface {
	GetAll(offset, limit int) []*models.Post
	GetById(id string) (*models.Post, error)
	CreatePost(userId string, post *dtos.CreatePostInput) (*models.Post, error)
	Delete(id string) error
}

type postService struct {
	postRepo repositories.PostRepository
}

func NewPostService(repo repositories.PostRepository) *postService {
	return &postService{postRepo: repo}
}

func (serv *postService) GetAll(offset, limit int) []*models.Post {
	return serv.postRepo.GetAll(offset, limit)
}

func (serv *postService) GetById(id string) (*models.Post, error) {
	return serv.postRepo.GetById(id)
}

func (serv *postService) CreatePost(userId string, post *dtos.CreatePostInput) (*models.Post, error) {
	_post := models.Post{
		Title:   post.Title,
		Content: post.Content,
		UserId:  userId,
	}
	createdPost := serv.postRepo.Create(&_post)
	return createdPost, nil
}

func (serv *postService) Delete(id string) error {
	return serv.postRepo.Delete(id)
}
