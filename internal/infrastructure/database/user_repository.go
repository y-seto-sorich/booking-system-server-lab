package database

import (
	"booking-system-server-lab/internal/adapter/gateway"
	"booking-system-server-lab/internal/infrastructure/database/entity"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// UserRepositoryの実装を返すコンストラクタ
func NewUserRepository(db *gorm.DB) gateway.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Save(user *entity.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindByID(id int) (*entity.User, error) {
	var user entity.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByEmail implements gateway.UserRepository.
func (r *userRepository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Delete(user *entity.User) error {
	return r.db.Delete(user).Error
}
