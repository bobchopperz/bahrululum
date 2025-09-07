package middleware

import (
	"net/http"

	"github.com/bobchopperz/bahrululum/internal/constants"
	"github.com/bobchopperz/bahrululum/internal/domain/service"
	"github.com/bobchopperz/bahrululum/internal/util"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func RequireRole(userService service.UserService, requiredRoles ...constants.Role) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userIDStr := c.Get("user_id")
			if userIDStr == nil {
				return util.ErrorResponse(c, http.StatusUnauthorized, "User not authenticated")
			}

			userID, ok := userIDStr.(uuid.UUID)
			if !ok {
				return util.ErrorResponse(c, http.StatusUnauthorized, "Invalid user ID")
			}

			user, err := userService.GetUser(c.Request().Context(), userID)
			if err != nil {
				return util.ErrorResponse(c, http.StatusUnauthorized, "User not found")
			}

			userRole, _ := constants.ParseRole(user.Role)
			hasRole := false
			for _, requiredRole := range requiredRoles {
				if userRole == requiredRole {
					hasRole = true
					break
				}
			}

			if !hasRole {
				return util.ErrorResponse(c, http.StatusForbidden, "Insufficient privileges")
			}

			c.Set("user_role", userRole.String())
			c.Set("user", user)
			return next(c)
		}
	}
}

func RequireAnyRole(userService service.UserService, roles ...constants.Role) echo.MiddlewareFunc {
	return RequireRole(userService, roles...)
}

func RequireAdmin(userService service.UserService) echo.MiddlewareFunc {
	return RequireRole(userService, constants.RoleAdmin)
}

func RequireMentor(userService service.UserService) echo.MiddlewareFunc {
	return RequireRole(userService, constants.RoleMentor)
}

func RequireMentorOrAdmin(userService service.UserService) echo.MiddlewareFunc {
	return RequireRole(userService, constants.RoleMentor, constants.RoleAdmin)
}

func RequireUser(userService service.UserService) echo.MiddlewareFunc {
	return RequireRole(userService, constants.RoleUser)
}

func RequireAnyRoleAccess(userService service.UserService) echo.MiddlewareFunc {
	return RequireRole(userService, constants.RoleUser, constants.RoleMentor, constants.RoleAdmin)
}
