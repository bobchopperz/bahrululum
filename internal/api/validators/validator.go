package validators

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func NewValidator() *CustomValidator {
	return &CustomValidator{validator: validator.New()}
}

func ValidationErrorResponse(c echo.Context, err error) error {
	return c.JSON(http.StatusBadRequest, map[string]interface{}{
		"error":   "Validation failed",
		"details": err.Error(),
	})
}
