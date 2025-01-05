package handler

import (
	"booking-system-server-lab/internal/adapter/controller/dto"
	"booking-system-server-lab/internal/adapter/presenter"
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

	var request dto.CreateUserRequest
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := u.userUsecase.CreateUser(ctx, request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	response := presenter.ToUserResponse(user)

	return c.JSON(http.StatusOK, response)
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
