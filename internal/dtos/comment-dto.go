package dtos

import "time"

type CommentInput struct {
	UserId  string
	PostId  string `json:"post_id" binding:"required"`
	Content string `json:"content" binding:"required"`
}
type CommentResponse struct {
	Id        string      `json:"id"`
	PostId    string      `json:"post_id"`
	Content   string      `json:"content"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	Author    UserReponse `json:"author"`
}
