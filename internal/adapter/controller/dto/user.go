package dto

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	BirthDay string `json:"birthDay" validate:"required"`
}
