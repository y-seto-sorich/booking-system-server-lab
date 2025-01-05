package application

import (
	"booking-system-server-lab/internal/adapter/controller/dto"
	"booking-system-server-lab/internal/adapter/gateway"
	"booking-system-server-lab/internal/application/usecase"
	"booking-system-server-lab/internal/domain/factory"
	"booking-system-server-lab/internal/domain/model"
	domain "booking-system-server-lab/internal/domain/service"
	"context"
	"errors"
)

type userService struct {
	userRepository gateway.UserRepository
	userFactory    factory.UserFactory
	domainService  domain.UserService
}

// UserUsecaseのインスタンスを作成するコンストラクタ
func NewUserService(userRepository gateway.UserRepository, userFactory factory.UserFactory, domainService domain.UserService) usecase.UserUsecase {
	return &userService{
		userRepository: userRepository,
		userFactory:    userFactory,
		domainService:  domainService,
	}
}

func (s *userService) CreateUser(ctx context.Context, request dto.CreateUserRequest) (*model.User, error) {
	// メールアドレスが既に登録されているかチェック
	exists, err := s.domainService.Exists(request.Email)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, errors.New("this email is already registered")
	}

	// ファクトリーを使ってユーザーを作成
	user, err := s.userFactory.CreateUserWithValidation(request)
	if err != nil {
		return nil, err
	}

	// ドメインモデル(model.User)をインフラ層のエンティティ(entity.User)に変換
	entityUser := gateway.ToEntityUser(user)

	// ユーザーをリポジトリに保存
	err = s.userRepository.Save(entityUser)
	if err != nil {
		return nil, err
	}

	// ドメインモデル(model.User)を戻り値として返す
	return user, nil
}

// DeleteUser implements usecase.UserUsecase.
func (s *userService) DeleteUser(id int) error {
	panic("unimplemented")
}

// GetUser implements usecase.UserUsecase.
func (s *userService) GetUser(id int) (*model.User, error) {
	panic("unimplemented")
}

// UpdateUser implements usecase.UserUsecase.
func (s *userService) UpdateUser(id int, name string, email string) (*model.User, error) {
	panic("unimplemented")
}
