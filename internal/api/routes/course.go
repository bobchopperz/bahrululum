package routes

import (
	"github.com/bobchopperz/bahrululum/internal/api/handlers"
	"github.com/bobchopperz/bahrululum/internal/domain/service"
	"github.com/labstack/echo/v4"
)

func SetupCoursesRoutes(e *echo.Echo, courseService service.CourseService) {
	courseHandler := handlers.NewCourseHandler(courseService)

	courses := e.Group("/api/courses")
	courses.GET("", courseHandler.GetCourses)
	courses.GET("/:id", courseHandler.GetCourse)
	courses.POST("", courseHandler.CreateCourse)
	courses.PUT("/:id", courseHandler.UpdateCourse)
	courses.DELETE("/:id", courseHandler.DeleteCourse)
}
