package handler

import echoLibrary "github.com/labstack/echo/v4"

type Handler struct {
	Echo Echo
}

type Echo struct {
	Handler    echoLibrary.HandlerFunc
	Middleware []echoLibrary.MiddlewareFunc
}

func CreateForEcho(h echoLibrary.HandlerFunc, middlewares ...echoLibrary.MiddlewareFunc) *Handler {
	return &Handler{
		Echo: Echo{
			Handler:    h,
			Middleware: middlewares,
		},
	}
}
