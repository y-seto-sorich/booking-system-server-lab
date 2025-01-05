package handler

import (
	"booking-system-server-lab/internal/application/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	CreateUser(c echo.Context) error
	GetUser(c echo.Context) error
	UpdateUser(c echo.Context) error
	DeleteUser(c echo.Context) error
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) UserHandler {
	return &userHandler{userUsecase}
}

// CreateUser implements UserHandler.
func (u *userHandler) CreateUser(c echo.Context) error {
	ctx := c.Request().Context()

	responce, err := u.userUsecase.CreateUser(ctx, "", "", "")
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, responce)
}

// GetUser implements UserHandler.
func (u *userHandler) GetUser(c echo.Context) error {
	panic("unimplemented")
}

// UpdateUser implements UserHandler.
func (u *userHandler) UpdateUser(c echo.Context) error {
	panic("unimplemented")
}

// DeleteUser implements UserHandler.
func (u *userHandler) DeleteUser(c echo.Context) error {
	panic("unimplemented")
}
