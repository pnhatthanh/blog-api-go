package dtos

import "time"

type CreatePostInput struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}
type PostResponse struct {
	Id          string            `json:"post_id"`
	Title       string            `json:"title"`
	Content     string            `json:"content"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
	UserReponse UserReponse       `json:"author"`
	Comments    []CommentResponse `json:"comments"`
}
