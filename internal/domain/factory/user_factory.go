package factory

import (
	"booking-system-server-lab/internal/domain/model"

	"github.com/google/uuid"
)

// Userエンティティを生成するファクトリーのインターフェース
type UserFactory interface {
	CreateUserWithValidation(name, email string, birthday string) (*model.User, error)
}

// UserFactoryインターフェースを実装する構造体
type userFactory struct{}

// UserFactoryのインスタンスを生成するコンストラクタ
func NewUserFactory() UserFactory {
	return &userFactory{}
}

// CreateUserWithValidation implements UserFactory.
func (u *userFactory) CreateUserWithValidation(name string, email string, birthday string) (*model.User, error) {

	// ユーザーIDを生成
	userID := uuid.New().String()

	// ドメインモデルのインスタンスを生成
	return model.NewUser(userID, name, email, birthday)
}
