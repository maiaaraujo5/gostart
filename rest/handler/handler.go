package handler

import (
	echoLibrary "github.com/labstack/echo/v4"
	"github.com/maiaaraujo5/gostart/rest/handler/wrapper"
)

type Handler struct {
	Echo Echo
}

type Echo struct {
	Handler    echoLibrary.HandlerFunc
	Middleware []echoLibrary.MiddlewareFunc
}

func CreateForEcho(h wrapper.Handler, middlewares ...echoLibrary.MiddlewareFunc) *Handler {
	return &Handler{
		Echo: Echo{
			Handler:    wrapper.Handle(h),
			Middleware: middlewares,
		},
	}
}
