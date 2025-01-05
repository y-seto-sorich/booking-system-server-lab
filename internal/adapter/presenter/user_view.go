package presenter

import (
	"booking-system-server-lab/internal/domain/model"
	"booking-system-server-lab/internal/util"
)

type UserResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	BirthDay string `json:"birthDay"`
}

func ToUserResponse(user *model.User) UserResponse {
	return UserResponse{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		BirthDay: util.ToDateString(user.BirthDay),
	}
}
