package errors

import (
	"github.com/labstack/echo/v4"
	"github.com/maiaaraujo5/gostart/errors"
	"github.com/maiaaraujo5/gostart/rest"
	"net/http"
)

func NewErrorResponse(c echo.Context, status int, message string) error {
	return c.JSON(status, rest.NewHttpError(status, message))
}

func ToErrorResponse(c echo.Context, err error) error {
	switch {
	case errors.IsNotFound(err):
		return c.JSON(http.StatusNotFound, rest.NewHttpError(http.StatusNotFound, err.Error()))
	case errors.IsBadRequest(err):
		return c.JSON(http.StatusBadRequest, rest.NewHttpError(http.StatusBadRequest, err.Error()))
	case errors.IsAlreadyExists(err):
		return c.JSON(http.StatusConflict, rest.NewHttpError(http.StatusConflict, err.Error()))
	default:
		return c.JSON(http.StatusInternalServerError, rest.NewHttpError(http.StatusInternalServerError, "internal server error"))
	}
}
