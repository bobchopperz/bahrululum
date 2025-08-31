package handlers

import (
	"net/http"
	"strconv"

	"github.com/bobchopperz/bahrululum/internal/domain/service"
	"github.com/bobchopperz/bahrululum/internal/util"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) GetUsers(c echo.Context) error {
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

	users, err := h.userService.GetUsers(c.Request().Context(), offset, limit)
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, "failed to fetch users")
	}

	return util.SuccessResponse(c, http.StatusOK, "Users retrieved successfully", map[string]interface{}{
		"users":  users,
		"offset": offset,
		"limit":  limit,
		"count":  len(users),
	})
}

func (h *UserHandler) GetUser(c echo.Context) error {
	idstr := c.Param("id")

	userID, err := uuid.Parse(idstr)
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve user")
	}

	user, err := h.userService.GetUser(c.Request().Context(), userID)
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve user")
	}

	return util.SuccessResponse(c, http.StatusOK, "User retrieved successfully", user)
}
