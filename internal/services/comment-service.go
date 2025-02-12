package services

import (
	"blogapi/internal/dtos"
	"blogapi/internal/models"
	"blogapi/internal/repositories"
	"errors"
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
func (serv *commentService) UpdateComment(id, userId string, comment *dtos.CommentInput) (*models.Comment, error) {
	_comment, err := serv.commentRepo.GetById(id)
	if err != nil {
		return nil, errors.New("Comment not found")
	}
	if userId != _comment.UserId {
		return nil, errors.New("Unauthorized: You can only update your own comment")
	}
	_comment.Content = comment.Content
	err = serv.commentRepo.UpdateComment(_comment)
	if err != nil {
		return nil, errors.New("Failed to update comment")
	}
	return _comment, nil
}

func (serv *commentService) Delete(id string) error {
	return serv.commentRepo.Delete(id)
}
