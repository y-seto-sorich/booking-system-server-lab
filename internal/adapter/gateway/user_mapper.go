package gateway

import (
	"booking-system-server-lab/internal/domain/model"
	"booking-system-server-lab/internal/infrastructure/database/entity"
)

func ToEntityUser(user *model.User) *entity.User {
	return &entity.User{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Birthday: user.Birthday,
	}
}
