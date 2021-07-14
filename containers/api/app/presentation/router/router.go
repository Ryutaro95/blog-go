package router

import (
	"net/http"

	"github.com/Ryutaro95/app/presentation/di"
	"github.com/labstack/echo"
)

func Router() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
	hc := di.InitializeHealthCheckHandler()
	e.GET("/health", hc.HealthCheck)

	return e
}
