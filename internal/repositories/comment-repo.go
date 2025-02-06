package repositories

import (
	"blogapi/internal/dtos"
	"blogapi/internal/models"

	"gorm.io/gorm"
)

type CommentRepository interface {
	GetAllComments(postId string, offset, limit int) []*models.Comment
	CreateComment(id string, comment *dtos.CommentInput) *models.Comment
	UpdateComment(comment *dtos.CommentInput) (*models.Comment, error)
	Delete(id string) error
}
type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *commentRepository {
	return &commentRepository{db: db}
}

func (repo *commentRepository) GetAllComments(postId string, offset, limit int) []*models.Comment {
	var comments []*models.Comment
	repo.db.Preload("User").Where("post_id= ?", postId).Order("created_at desc").
		Limit(limit).Offset(offset).Find(&comments)
	return comments
}

func (repo *commentRepository) CreateComment(comment *dtos.CommentInput) *models.Comment {
	_comment := models.Comment{
		Content: comment.Content,
		PostId:  comment.PostId,
		UserId:  comment.UserId,
	}
	repo.db.Create(&_comment)
	return &_comment
}

func (repo *commentRepository) UpdateComment(id string, comment *dtos.CommentInput) (*models.Comment, error) {
	var _comment models.Comment
	if err := repo.db.First(&_comment, "id = ?", id).Error; err != nil {
		return nil, err
	}
	_comment.Content = comment.Content
	if err := repo.db.Save(&_comment).Error; err != nil {
		return nil, err
	}
	return &_comment, nil
}

func (repo *commentRepository) Delete(id string) error {
	var _comment models.Comment
	if err := repo.db.First(&_comment, "id = ?", id).Error; err != nil {
		return err
	}
	if err := repo.db.Delete(&_comment).Error; err != nil {
		return err
	}
	return nil
}
