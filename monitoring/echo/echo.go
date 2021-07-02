package echo

import (
	sentryecho "github.com/getsentry/sentry-go/echo"
	"github.com/labstack/echo/v4"
	"github.com/maiaaraujo5/gostart/log/logger"
	"github.com/maiaaraujo5/gostart/monitoring/sentry"
	"log"
)

func Middleware() echo.MiddlewareFunc {
	err := sentry.Init()
	if err != nil {
		log.Fatalln("error to initialize sentry")
	}
	logger.Info("sentry for echo initialized successfully")
	return sentryecho.New(sentryecho.Options{
		Repanic: true,
	})
}
