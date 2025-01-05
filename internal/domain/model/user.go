package model

import (
	"booking-system-server-lab/internal/adapter/controller/dto"
	"errors"
	"time"
)

type User struct {
	ID       string
	Name     string
	Email    string
	BirthDay time.Time
}

func NewUser(id string, request dto.CreateUserRequest) (*User, error) {
	if request.Name == "" || request.Email == "" {
		return nil, errors.New("name and email are required")
	}

	// if !isValidEmail(email) {
	// 	return nil, errors.New("invalid email format")
	// }

	birthdayTime, err := time.Parse("2006-01-02", request.BirthDay)
	if err != nil {
		return nil, errors.New("invalid date format for birthday")
	}

	return &User{
		ID:       id,
		Name:     request.Name,
		Email:    request.Email,
		BirthDay: birthdayTime,
	}, nil
}

func (u *User) UpdateName(newName string) error {
	if newName == "" {
		return errors.New("new name cannot be empty")
	}
	u.Name = newName
	return nil
}
