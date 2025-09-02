package middleware

import (
	"net/http"
	"strings"

	"github.com/bobchopperz/bahrululum/internal/domain/service"
	"github.com/bobchopperz/bahrululum/internal/util"
	"github.com/labstack/echo/v4"
)

func JWTAuth(authService service.AuthService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return util.ErrorResponse(c, http.StatusUnauthorized, "Missing authentication header")
			}

			tokenParts := strings.SplitN(authHeader, " ", 2)
			if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
				return util.ErrorResponse(c, http.StatusUnauthorized, "Invalid authorization header format")
			}

			tokenString := tokenParts[1]
			if tokenString == "" {
				return util.ErrorResponse(c, http.StatusUnauthorized, "Missing token")
			}

			claims, err := authService.ValidateToken(tokenString)
			if err != nil {
				return util.ErrorResponse(c, http.StatusUnauthorized, "Invalid or expired token")
			}

			c.Set("user_id", claims.UserID)
			c.Set("user_claims", claims)

			return nil
		}
	}
}
