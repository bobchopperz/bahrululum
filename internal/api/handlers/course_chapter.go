package handlers

import (
	"net/http"
	"strconv"

	"github.com/bobchopperz/bahrululum/internal/api/validators"
	"github.com/bobchopperz/bahrululum/internal/domain/models"
	"github.com/bobchopperz/bahrululum/internal/domain/service"
	"github.com/bobchopperz/bahrululum/internal/util"
	"github.com/labstack/echo/v4"
)

type CourseChapterHandler struct {
	chapterService service.CourseChapterService
}

func NewCourseChapterHandler(chapterService service.CourseChapterService) *CourseChapterHandler {
	return &CourseChapterHandler{chapterService: chapterService}
}

func (h *CourseChapterHandler) GetChapters(c echo.Context) error {
	courseIDStr := c.QueryParam("course_id")
	if courseIDStr == "" {
		return util.ErrorResponse(c, http.StatusBadRequest, "course_id parameter is required")
	}

	courseID, err := strconv.ParseUint(courseIDStr, 10, 32)
	if err != nil {
		return util.ErrorResponse(c, http.StatusBadRequest, "Invalid course_id parameter")
	}

	chapters, err := h.chapterService.GetChaptersByCourse(c.Request().Context(), uint(courseID))
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve chapters")
	}

	return util.SuccessResponse(c, http.StatusOK, "Chapters retrieved successfully", map[string]interface{}{
		"chapters": chapters,
		"count":    len(chapters),
	})
}

func (h *CourseChapterHandler) GetChapter(c echo.Context) error {
	idStr := c.Param("id")

	chapterID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return util.ErrorResponse(c, http.StatusBadRequest, "Invalid chapter ID")
	}

	chapter, err := h.chapterService.GetChapter(c.Request().Context(), uint(chapterID))
	if err != nil {
		return util.ErrorResponse(c, http.StatusNotFound, "Chapter not found")
	}

	return util.SuccessResponse(c, http.StatusOK, "Chapter retrieved successfully", chapter)
}

func (h *CourseChapterHandler) CreateChapter(c echo.Context) error {
	var req models.CreateCourseChapterRequest

	if err := c.Bind(&req); err != nil {
		return util.ErrorResponse(c, http.StatusBadRequest, "Invalid request body")
	}

	if err := c.Validate(&req); err != nil {
		return validators.ValidationErrorResponse(c, err)
	}

	chapter, err := h.chapterService.CreateChapter(c.Request().Context(), &req)
	if err != nil {
		return util.ErrorResponse(c, http.StatusUnprocessableEntity, "Failed to create chapter")
	}

	return util.SuccessResponse(c, http.StatusCreated, "Chapter created successfully", chapter)
}

func (h *CourseChapterHandler) UpdateChapter(c echo.Context) error {
	idStr := c.Param("id")

	chapterID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return util.ErrorResponse(c, http.StatusBadRequest, "Invalid chapter ID")
	}

	var req models.UpdateCourseChapterRequest
	if err := c.Bind(&req); err != nil {
		return util.ErrorResponse(c, http.StatusBadRequest, "Invalid request body")
	}

	if err := c.Validate(&req); err != nil {
		return validators.ValidationErrorResponse(c, err)
	}

	chapter, err := h.chapterService.UpdateChapter(c.Request().Context(), uint(chapterID), &req)
	if err != nil {
		return util.ErrorResponse(c, http.StatusUnprocessableEntity, "Failed to update chapter")
	}

	return util.SuccessResponse(c, http.StatusOK, "Chapter updated successfully", chapter)
}

func (h *CourseChapterHandler) DeleteChapter(c echo.Context) error {
	idStr := c.Param("id")

	chapterID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return util.ErrorResponse(c, http.StatusBadRequest, "Invalid chapter ID")
	}

	err = h.chapterService.DeleteChapter(c.Request().Context(), uint(chapterID))
	if err != nil {
		return util.ErrorResponse(c, http.StatusUnprocessableEntity, "Failed to delete chapter")
	}

	return util.SuccessResponse(c, http.StatusOK, "Chapter deleted successfully", nil)
}

func (h *CourseChapterHandler) GetChapterWithContents(c echo.Context) error {
	idStr := c.Param("id")

	chapterID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return util.ErrorResponse(c, http.StatusBadRequest, "Invalid chapter ID")
	}

	chapter, err := h.chapterService.GetChapterWithContents(c.Request().Context(), uint(chapterID))
	if err != nil {
		return util.ErrorResponse(c, http.StatusNotFound, "Chapter not found")
	}

	return util.SuccessResponse(c, http.StatusOK, "Chapter with contents retrieved successfully", chapter)
}
