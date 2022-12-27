package health

import (
	"context"
	"fmt"
	echoLibrary "github.com/labstack/echo/v4"
	"github.com/maiaaraujo5/gostart/echo"
	"github.com/maiaaraujo5/gostart/echo/plugins/health/response"
	"github.com/maiaaraujo5/gostart/health"
	"github.com/maiaaraujo5/gostart/log/logger"
	"net/http"
)

func NewHealth(ctx context.Context) echo.Plugin {
	logger.Trace("registering health echo plugin")

	c, err := NewConfig()
	if err != nil {
		logger.Fatal("error to get config from health echo plugin")
	}

	return func(ctx context.Context, client *echoLibrary.Echo) error {

		client.GET(c.Route, handler)

		logger.Debug(fmt.Sprintf("health router configured on %s in echo", c.Route))

		return nil
	}
}

func handler(c echoLibrary.Context) error {
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	healths := health.Check(ctx)

	if len(healths) == 0 {
		return c.JSON(http.StatusOK, &response.Response{
			Status: "Healthy",
		})
	}

	var responses []response.Response
	for _, h := range healths {
		responses = append(responses, response.Response{
			Name:        h.Health.Name,
			Description: h.Health.Description,
			Status: func() string {
				if h.Err != nil {
					return "Unhealthy"
				}
				return "Healthy"
			}(),
			Error: h.Err,
		})
	}

	return c.JSON(http.StatusOK, responses)
}
