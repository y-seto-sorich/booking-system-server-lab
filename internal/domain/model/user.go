package model

import (
	"errors"
	"time"
)

type User struct {
	ID       string
	Name     string
	Email    string
	Birthday time.Time
}

func NewUser(id string, name string, email string, birthday string) (*User, error) {
	if name == "" || email == "" {
		return nil, errors.New("name and email are required")
	}

	// if !isValidEmail(email) {
	// 	return nil, errors.New("invalid email format")
	// }

	birthdayTime, err := time.Parse("2006-01-02", birthday)
	if err != nil {
		return nil, errors.New("invalid date format for birthday")
	}

	return &User{
		ID:       id,
		Name:     name,
		Email:    email,
		Birthday: birthdayTime,
	}, nil
}

func (u *User) UpdateName(newName string) error {
	if newName == "" {
		return errors.New("new name cannot be empty")
	}
	u.Name = newName
	return nil
}
