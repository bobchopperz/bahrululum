package routes

import (
	"github.com/bobchopperz/bahrululum/internal/api/handlers"
	"github.com/labstack/echo/v4"
)

func SetupHealthRoutes(e *echo.Echo) {
	e.GET("/api/health", handlers.HealthCheck)
}
