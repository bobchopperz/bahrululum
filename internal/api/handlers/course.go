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

type CourseHandler struct {
	courseService service.CourseService
}

func NewCourseHandler(courseService service.CourseService) *CourseHandler {
	return &CourseHandler{courseService: courseService}
}

func (h *CourseHandler) GetCourses(c echo.Context) error {
	offsetStr := c.QueryParam("offset")
	limitStr := c.QueryParam("limit")

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 || limit > 100 {
		limit = 10
	}

	courses, err := h.courseService.GetCourses(c.Request().Context(), offset, limit)
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, "failed to fetch courses")
	}

	return util.SuccessResponse(c, http.StatusOK, "Courses retrieved successfully", map[string]interface{}{
		"courses": courses,
		"offset":  offset,
		"limit":   limit,
		"count":   len(courses),
	})
}

func (h *CourseHandler) GetCourse(c echo.Context) error {
	idstr := c.Param("id")

	courseID, err := strconv.ParseUint(idstr, 10, 32)
	if err != nil {
		return util.ErrorResponse(c, http.StatusBadRequest, "Invalid course ID")
	}

	course, err := h.courseService.GetCourse(c.Request().Context(), uint(courseID))
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve course")
	}

	return util.SuccessResponse(c, http.StatusOK, "Course retrieved successfully", course)
}

func (h *CourseHandler) CreateCourse(c echo.Context) error {
	var req models.CreateCourseRequest

	if err := c.Bind(&req); err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, "Invalid request")
	}

	if err := c.Validate(&req); err != nil {
		return validators.ValidationErrorResponse(c, err)
	}

	course, err := h.courseService.CreateCourse(c.Request().Context(), &req)
	if err != nil {
		return util.ErrorResponse(c, http.StatusUnprocessableEntity, "Something went wrong")
	}

	return util.SuccessResponse(c, http.StatusCreated, "Course created successfully", course)
}

func (h *CourseHandler) UpdateCourse(c echo.Context) error {
	idStr := c.Param("id")
	courseID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return util.ErrorResponse(c, http.StatusBadRequest, "Invalid course ID")
	}

	var req models.CreateCourseRequest
	if err := c.Bind(&req); err != nil {
		return util.ErrorResponse(c, http.StatusBadRequest, "Invalid request")
	}

	if err := c.Validate(&req); err != nil {
		return validators.ValidationErrorResponse(c, err)
	}

	updates := map[string]interface{}{
		"name":        req.Name,
		"description": req.Description,
	}

	course, err := h.courseService.UpdateCourse(c.Request().Context(), uint(courseID), updates)
	if err != nil {
		return util.ErrorResponse(c, http.StatusUnprocessableEntity, "Failed to update course")
	}

	return util.SuccessResponse(c, http.StatusOK, "Course updated successfully", course)
}

func (h *CourseHandler) DeleteCourse(c echo.Context) error {
	idStr := c.Param("id")
	courseID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return util.ErrorResponse(c, http.StatusBadRequest, "Invalid course ID")
	}

	err = h.courseService.DeleteCourse(c.Request().Context(), uint(courseID))
	if err != nil {
		return util.ErrorResponse(c, http.StatusUnprocessableEntity, "Failed to delete course")
	}

	return util.SuccessResponse(c, http.StatusOK, "Course deleted successfully", nil)
}
