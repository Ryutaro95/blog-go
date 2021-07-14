package handler

import (
	"net/http"

	"github.com/Ryutaro95/app/application/usecase"
	"github.com/Ryutaro95/app/infrastructure/logger"
	"github.com/labstack/echo"
)

type HealthCheckHandler struct {
	healthCheckUsecaseAccessor usecase.HealthCheckUsecaseAccessor
}

func NewHealthCheckHandler(hc usecase.HealthCheckUsecaseAccessor) *HealthCheckHandler {
	return &HealthCheckHandler{
		healthCheckUsecaseAccessor: hc,
	}
}

func (hc HealthCheckHandler) HealthCheck(c echo.Context) error {
	logger.Info("start health check")

	if err := hc.healthCheckUsecaseAccessor.HealthCheck(); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"status": "failed", "message": err.Error()})
	}

	logger.Info("successful health check")

	return c.JSON(http.StatusOK, nil)
}
