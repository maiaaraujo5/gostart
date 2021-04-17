package rest

import "github.com/maiaaraujo5/gostart/rest/handler"

type Rest interface {
	AddRoute(method, route string, handler handler.Handler)
	Start()
}
