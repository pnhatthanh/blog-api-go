package dtos

type CommentInput struct {
	UserId  string
	PostId  string `json:"post_id" binding:"required"`
	Content string `json:"content" binding:"required"`
}
