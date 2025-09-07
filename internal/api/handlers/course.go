package handlers

import (
	"net/http"
	"strconv"

	"github.com/bobchopperz/bahrululum/internal/domain/service"
	"github.com/bobchopperz/bahrululum/internal/util"
	"github.com/google/uuid"
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

	courseID, err := uuid.Parse(idstr)
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve course")
	}

	course, err := h.courseService.GetCourse(c.Request().Context(), courseID)
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve course")
	}

	return util.SuccessResponse(c, http.StatusOK, "Course retrieved successfully", course)
}
