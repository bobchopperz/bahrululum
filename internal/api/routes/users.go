package routes

import (
	"github.com/bobchopperz/bahrululum/internal/api/handlers"
	"github.com/bobchopperz/bahrululum/internal/domain/service"
	"github.com/labstack/echo/v4"
)

func SetupUsersRoutes(e *echo.Echo, userService service.UserService) {
	userHandler := handlers.NewUserHandler(userService)

	users := e.Group("/api/users")
	users.GET("", userHandler.GetUsers)
	users.GET("/:id", userHandler.GetUser)
}
