package wrapper

import (
	sentryecho "github.com/getsentry/sentry-go/echo"
	echoLibrary "github.com/labstack/echo/v4"
	"github.com/maiaaraujo5/gostart/rest/echo/config"
	"github.com/maiaaraujo5/gostart/rest/echo/errors"
)

type (
	Handler func(c echoLibrary.Context) error
)

func Handle(handler Handler) echoLibrary.HandlerFunc {
	return func(context echoLibrary.Context) error {
		err := handler(context)

		if config.SentryEnabled() {
			hub := sentryecho.GetHubFromContext(context)
			if err != nil {
				hub.CaptureException(err)
			}
		}

		if err != nil {
			return errors.ToErrorResponse(context, err)
		}

		return err
	}
}
