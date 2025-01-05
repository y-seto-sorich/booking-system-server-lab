package usecase

import (
	"booking-system-server-lab/internal/adapter/controller/dto"
	"booking-system-server-lab/internal/domain/model"
	"context"
)

// ユーザーに関連するユースケースを提供するインターフェース
type UserUsecase interface {
	CreateUser(ctx context.Context, request dto.CreateUserRequest) (*model.User, error)
	UpdateUser(id int, name string, email string) (*model.User, error)
	GetUser(id int) (*model.User, error)
	DeleteUser(id int) error
}
