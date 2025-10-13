package routes

import (
	"github.com/bobchopperz/bahrululum/internal/api/handlers"
	"github.com/bobchopperz/bahrululum/internal/api/middleware"
	"github.com/bobchopperz/bahrululum/internal/domain/service"
	"github.com/labstack/echo/v4"
)

func SetupEnrollmentRoutes(e *echo.Echo, enrollmentService service.EnrollmentService, authService service.AuthService) {
	h := handlers.NewEnrollmentHandler(enrollmentService)

	enrollments := e.Group("/api/enrollments")
	enrollments.Use(middleware.JWTAuth(authService))

	enrollments.GET("/my", h.GetMyEnrollments)
	enrollments.GET("/:course_id", h.GetEnrollment)
	enrollments.POST("", h.Create)
}
