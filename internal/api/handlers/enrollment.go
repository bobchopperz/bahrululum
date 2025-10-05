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

type EnrollmentHandler struct {
	enrollmentService service.EnrollmentService
}

func NewEnrollmentHandler(s service.EnrollmentService) *EnrollmentHandler {
	return &EnrollmentHandler{enrollmentService: s}
}

func (h *EnrollmentHandler) Create(c echo.Context) error {
	var req models.CreateEnrollmentRequest

	if err := c.Bind(&req); err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, "Invalid request")
	}

	if err := c.Validate(&req); err != nil {
		return validators.ValidationErrorResponse(c, err)
	}

	entity, err := h.enrollmentService.Create(c.Request().Context(), &req)
	if err != nil {
		return util.ErrorResponse(c, http.StatusUnprocessableEntity, "Something went wrong")
	}

	return util.SuccessResponse(c, http.StatusCreated, "Enrollment retrieved successfully", entity)
}

func (h *EnrollmentHandler) GetEnrollment(c echo.Context) error {
	strCourseID := c.Param("course_id")

	courseID, err := strconv.ParseUint(strCourseID, 10, 32)
	if err != nil {
		return util.ErrorResponse(c, http.StatusBadRequest, "Invalid course ID")
	}

	entity, err := h.enrollmentService.GetByCouseID(c.Request().Context(), uint(courseID))
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve enrollment")
	}

	return util.SuccessResponse(c, http.StatusOK, "Course retrieved successfully", entity)
}
