package routes

import (
	"github.com/bobchopperz/bahrululum/internal/api/handlers"
	"github.com/bobchopperz/bahrululum/internal/domain/service"
	"github.com/labstack/echo/v4"
)

type AuthRoutesOpts struct {
	AuthService service.AuthService
	UserService service.UserService
}

func SetupAuthRoutes(e *echo.Echo, opts AuthRoutesOpts) {
	authHandler := handlers.NewAuthHandler(opts.AuthService, opts.UserService)

	e.POST("/api/login", authHandler.Login)
	e.POST("/api/register", authHandler.Register)
}
