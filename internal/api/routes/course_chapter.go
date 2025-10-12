package routes

import (
	"github.com/bobchopperz/bahrululum/internal/api/handlers"
	"github.com/bobchopperz/bahrululum/internal/domain/service"
	"github.com/labstack/echo/v4"
)

func SetupCourseChapterRoutes(e *echo.Echo, chapterService service.CourseChapterService) {
	chapterHandler := handlers.NewCourseChapterHandler(chapterService)

	chapters := e.Group("/api/chapters")
	chapters.GET("", chapterHandler.GetChapters)
	chapters.GET("/:id", chapterHandler.GetChapter)
	chapters.GET("/:id/contents", chapterHandler.GetChapterWithContents)
	chapters.POST("", chapterHandler.CreateChapter)
	chapters.PUT("/:id", chapterHandler.UpdateChapter)
	chapters.DELETE("/:id", chapterHandler.DeleteChapter)
}
