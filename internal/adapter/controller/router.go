package controller

import (
	"booking-system-server-lab/internal/adapter/controller/handler"
	"booking-system-server-lab/internal/adapter/controller/middleware"

	"github.com/labstack/echo/v4"
)

type HandlersDependency struct {
	UserHandler handler.UserHandler
}

func SetupRoutes(e *echo.Echo, handlersDependency *HandlersDependency) {
	// ミドルウェアの設定
	e.Use(middleware.Logger())

	e.POST("/users", handlersDependency.UserHandler.CreateUser)
	e.GET("/users/:id", handlersDependency.UserHandler.GetUser)
	e.PUT("/users/:id", handlersDependency.UserHandler.UpdateUser)
	e.DELETE("/users/:id", handlersDependency.UserHandler.DeleteUser)
}
