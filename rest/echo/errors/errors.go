package errors

import (
	"github.com/labstack/echo/v4"
	"github.com/maiaaraujo5/gostart/errors"
	"github.com/maiaaraujo5/gostart/rest/httperror"
	"net/http"
)

func NewErrorResponse(c echo.Context, status int, message string) error {
	return c.JSON(status, httperror.New(status, message))
}

func ToErrorResponse(c echo.Context, err error) error {
	switch {
	case errors.IsNotFound(err):
		return c.JSON(http.StatusNotFound, httperror.New(http.StatusNotFound, err.Error()))
	case errors.IsBadRequest(err):
		return c.JSON(http.StatusBadRequest, httperror.New(http.StatusBadRequest, err.Error()))
	case errors.IsAlreadyExists(err):
		return c.JSON(http.StatusConflict, httperror.New(http.StatusConflict, err.Error()))
	default:
		return c.JSON(http.StatusInternalServerError, httperror.New(http.StatusInternalServerError, "internal server error"))
	}
}
