package gateway

import (
	"booking-system-server-lab/internal/infrastructure/database/entity"
)

// ユーザー情報の永続化を行うインターフェース
type UserRepository interface {
	Save(user *entity.User) error
	FindByID(id int) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	Delete(user *entity.User) error
}
