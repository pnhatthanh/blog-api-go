package dtos

type CreatePostInput struct {
	UserId  string
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}
