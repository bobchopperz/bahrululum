package handlers

import (
	"net/http"

	"github.com/bobchopperz/bahrululum/internal/api/validators"
	"github.com/bobchopperz/bahrululum/internal/constants"
	"github.com/bobchopperz/bahrululum/internal/domain/models"
	"github.com/bobchopperz/bahrululum/internal/domain/service"
	"github.com/bobchopperz/bahrululum/internal/util"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authService service.AuthService
	userService service.UserService
}

func NewAuthHandler(authService service.AuthService, userService service.UserService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		userService: userService,
	}
}

func (h *AuthHandler) Register(c echo.Context) error {
	var req models.CreateUserRequest

	if err := c.Bind(&req); err != nil {
		return util.ErrorResponse(c, http.StatusBadRequest, "Invalid request body")
	}

	req.Role = constants.RoleUser.String()

	if err := c.Validate(&req); err != nil {
		return validators.ValidationErrorResponse(c, err)
	}

	user, err := h.userService.CreateUser(c.Request().Context(), &req)
	if err != nil {
		return err
	}

	// Generate tokens for the new user
	tokens, err := h.authService.GenerateToken(user.ID)
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, "Failed to generate token")
	}

	tokens.User = user

	return util.SuccessResponse(c, http.StatusCreated, "User registered successfully", tokens)
}

func (h *AuthHandler) Login(c echo.Context) error {
	var req models.LoginRequest

	if err := c.Bind(&req); err != nil {
		return util.ErrorResponse(c, http.StatusBadRequest, "Invalid request body")
	}

	if err := c.Validate(&req); err != nil {
		return validators.ValidationErrorResponse(c, err)
	}

	tokens, err := h.authService.Login(c.Request().Context(), &req)
	if err != nil {
		return util.ErrorResponse(c, http.StatusUnauthorized, err.Error())
	}

	return util.SuccessResponse(c, http.StatusOK, "Login successful", tokens)
}
