package routes

import (
	"github.com/bobchopperz/bahrululum/internal/api/handlers"
	"github.com/bobchopperz/bahrululum/internal/domain/service"
	"github.com/labstack/echo/v4"
)

func SetupCourseContentRoutes(e *echo.Echo, contentService service.CourseContentService) {
	contentHandler := handlers.NewCourseContentHandler(contentService)

	contents := e.Group("/api/contents")
	contents.GET("", contentHandler.GetContents)
	contents.GET("/:id", contentHandler.GetContent)
	contents.POST("", contentHandler.CreateContent)
	contents.PUT("/:id", contentHandler.UpdateContent)
	contents.DELETE("/:id", contentHandler.DeleteContent)
}
