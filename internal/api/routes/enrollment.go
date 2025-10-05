package routes

import (
	"github.com/bobchopperz/bahrululum/internal/api/handlers"
	"github.com/bobchopperz/bahrululum/internal/domain/service"
	"github.com/labstack/echo/v4"
)

func SetupEnrollmentRoutes(e *echo.Echo, s service.EnrollmentService) {
	h := handlers.NewEnrollmentHandler(s)

	courses := e.Group("/api/enrollments")
	courses.GET("/:course_id", h.GetEnrollment)
	courses.POST("", h.Create)
}
