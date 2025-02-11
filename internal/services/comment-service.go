package services

import (
	"blogapi/internal/dtos"
	"blogapi/internal/models"
	"blogapi/internal/repositories"
)

type CommentService interface {
	GetAllComment(postId string, offset, limit int) []*models.Comment
	CreateComment(userId string, comment *dtos.CommentInput) *models.Comment
	UpdateComment(id, userId string, comment *dtos.CommentInput) (*models.Comment, error)
	Delete(id string) error
}

type commentService struct {
	commentRepo repositories.CommentRepository
}

func NewCommentService(repo repositories.CommentRepository) *commentService {
	return &commentService{
		commentRepo: repo,
	}
}
func (serv *commentService) GetAllComment(postId string, offset, limit int) []*models.Comment {
	return serv.commentRepo.GetAllComments(postId, offset, limit)
}
func (serv *commentService) CreateComment(userId string, comment *dtos.CommentInput) *models.Comment {
	_comment := models.Comment{
		Content: comment.Content,
		PostId:  comment.PostId,
		UserId:  userId,
	}
	return serv.commentRepo.CreateComment(&_comment)
}
func (serv *commentService) UpdateComment(id string, comment *dtos.CommentInput) (*models.Comment, error) {
	_comment, err := serv.commentRepo.GetById(id)
	if err != nil {
		return nil, err
	}
	_comment.Content = comment.Content
	return _comment, serv.commentRepo.UpdateComment(_comment)
}
func (serv *commentService) Delete(id string) error {
	return serv.commentRepo.Delete(id)
}
