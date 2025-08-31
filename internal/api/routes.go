package api

import (
	"github.com/bobchopperz/bahrululum/internal/api/handlers"
	"github.com/bobchopperz/bahrululum/internal/domain/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupRoutes(e *echo.Echo, userService service.UserService) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())

	userHandler := handlers.NewUserHandler(userService)

	users := e.Group("/users")
	users.GET("", userHandler.GetUsers)
	users.GET("/:id", userHandler.GetUser)
}
