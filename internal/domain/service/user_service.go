package domain

import (
	"booking-system-server-lab/internal/adapter/gateway"
)

type UserService interface {
	Exists(email string) (bool, error)
}

// UserServiceインターフェースを実装する構造体
type userService struct {
	userRepository gateway.UserRepository
}

// UserServiceのインスタンスを返すコンストラクタ
func NewUserService(userRepository gateway.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) Exists(email string) (bool, error) {
	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	if user != nil {
		return true, nil // 存在すればtrueを返す
	}
	return false, nil // 存在しなければfalseを返す
}
