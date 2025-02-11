package dtos

type CreateUserInput struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}
type UpdateUserInput struct {
	FullName string `json:"full_name" `
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"min=8"`
}
type UserLogin struct{
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}
type UserReponse struct {
	UserId string `json:"user_id"`
	FullName string `json:"full_name" `
	Email    string `json:"email"`
}

