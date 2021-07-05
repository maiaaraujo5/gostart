package wrapper

import (
	"github.com/getsentry/sentry-go"
	sentryecho "github.com/getsentry/sentry-go/echo"
	"github.com/google/uuid"
	echoLibrary "github.com/labstack/echo/v4"
	"github.com/maiaaraujo5/gostart/rest/echo/config"
	"github.com/maiaaraujo5/gostart/rest/echo/errors"
)

type (
	Handler func(c echoLibrary.Context) error
)

func Handle(handler Handler) echoLibrary.HandlerFunc {
	return func(context echoLibrary.Context) error {
		span := sentry.StartSpan(context.Request().Context(), "request", sentry.TransactionName(uuid.New().String()))
		err := handler(context)

		if config.SentryEnabled() {
			hub := sentryecho.GetHubFromContext(context)
			if err != nil {
				hub.CaptureException(err)
			}
		}

		if err != nil {
			span.Finish()
			return errors.ToErrorResponse(context, err)
		}

		span.Finish()
		return err
	}
}
