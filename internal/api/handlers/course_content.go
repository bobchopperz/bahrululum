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

type CourseContentHandler struct {
	contentService service.CourseContentService
}

func NewCourseContentHandler(contentService service.CourseContentService) *CourseContentHandler {
	return &CourseContentHandler{contentService: contentService}
}

func (h *CourseContentHandler) GetContents(c echo.Context) error {
	chapterIDStr := c.QueryParam("chapter_id")
	if chapterIDStr == "" {
		return util.ErrorResponse(c, http.StatusBadRequest, "chapter_id parameter is required")
	}

	chapterID, err := strconv.ParseUint(chapterIDStr, 10, 32)
	if err != nil {
		return util.ErrorResponse(c, http.StatusBadRequest, "Invalid chapter_id parameter")
	}

	contents, err := h.contentService.GetContentsByChapter(c.Request().Context(), uint(chapterID))
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve contents")
	}

	return util.SuccessResponse(c, http.StatusOK, "Contents retrieved successfully", map[string]interface{}{
		"contents": contents,
		"count":    len(contents),
	})
}

func (h *CourseContentHandler) GetContent(c echo.Context) error {
	idStr := c.Param("id")

	contentID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return util.ErrorResponse(c, http.StatusBadRequest, "Invalid content ID")
	}

	content, err := h.contentService.GetContent(c.Request().Context(), uint(contentID))
	if err != nil {
		return util.ErrorResponse(c, http.StatusNotFound, "Content not found")
	}

	return util.SuccessResponse(c, http.StatusOK, "Content retrieved successfully", content)
}

func (h *CourseContentHandler) CreateContent(c echo.Context) error {
	var req models.CreateCourseContentRequest

	if err := c.Bind(&req); err != nil {
		return util.ErrorResponse(c, http.StatusBadRequest, "Invalid request body")
	}

	if err := c.Validate(&req); err != nil {
		return validators.ValidationErrorResponse(c, err)
	}

	content, err := h.contentService.CreateContent(c.Request().Context(), &req)
	if err != nil {
		return util.ErrorResponse(c, http.StatusUnprocessableEntity, "Failed to create content")
	}

	return util.SuccessResponse(c, http.StatusCreated, "Content created successfully", content)
}

func (h *CourseContentHandler) UpdateContent(c echo.Context) error {
	idStr := c.Param("id")

	contentID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return util.ErrorResponse(c, http.StatusBadRequest, "Invalid content ID")
	}

	var req models.UpdateCourseContentRequest
	if err := c.Bind(&req); err != nil {
		return util.ErrorResponse(c, http.StatusBadRequest, "Invalid request body")
	}

	if err := c.Validate(&req); err != nil {
		return validators.ValidationErrorResponse(c, err)
	}

	content, err := h.contentService.UpdateContent(c.Request().Context(), uint(contentID), &req)
	if err != nil {
		return util.ErrorResponse(c, http.StatusUnprocessableEntity, "Failed to update content")
	}

	return util.SuccessResponse(c, http.StatusOK, "Content updated successfully", content)
}

func (h *CourseContentHandler) DeleteContent(c echo.Context) error {
	idStr := c.Param("id")

	contentID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return util.ErrorResponse(c, http.StatusBadRequest, "Invalid content ID")
	}

	err = h.contentService.DeleteContent(c.Request().Context(), uint(contentID))
	if err != nil {
		return util.ErrorResponse(c, http.StatusUnprocessableEntity, "Failed to delete content")
	}

	return util.SuccessResponse(c, http.StatusOK, "Content deleted successfully", nil)
}
