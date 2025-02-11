package repositories

import (
	"blogapi/internal/models"
	"gorm.io/gorm"
)

type CommentRepository interface {
	GetAllComments(postId string, offset, limit int) []*models.Comment
	CreateComment(comment *models.Comment) *models.Comment
	UpdateComment( comment *models.Comment) error
	Delete(id string) error
	GetById(id string) (*models.Comment,error)
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

func (repo *commentRepository) CreateComment(comment *models.Comment) *models.Comment {
	repo.db.Create(comment)
	return comment
}

func(repo *commentRepository) GetById(id string)(*models.Comment,error){
	var comment models.Comment
	if err:=repo.db.First(&comment,"id=?",id).Error; err!=nil{
		return nil,err
	}
	return &comment,nil
}

func (repo *commentRepository) UpdateComment(comment *models.Comment) error {
	if err := repo.db.Save(&comment).Error; err != nil {
		return err
	}
	return  nil
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
