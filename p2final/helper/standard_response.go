package helper

import (
	"p2final/dto"

	"github.com/labstack/echo/v4"
)

func SendError(c echo.Context, status int, code, message string, details interface{}) error {
	return c.JSON(status, dto.ErrorResponse{
		Status:  status,
		Code:    code,
		Message: message,
		Details: details,
	})
}

func SendSuccess(c echo.Context, status int, message string, data interface{}) error {
	return c.JSON(status, dto.SuccessResponse{
		Status:  status,
		Message: message,
		Data:    data,
	})
}
