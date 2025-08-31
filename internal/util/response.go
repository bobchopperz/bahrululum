package util

import "github.com/labstack/echo/v4"

type Reponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error"`
}

func SuccessResponse(c echo.Context, statusCode int, message string, data interface{}) error {
	return c.JSON(statusCode, Reponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func ErrorResponse(c echo.Context, statusCode int, message string) error {
	return c.JSON(statusCode, Reponse{
		Success: false,
		Message: "Something went wrong",
		Error:   message,
	})
}
