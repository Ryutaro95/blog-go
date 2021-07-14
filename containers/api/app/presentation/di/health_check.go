package di

import (
	"github.com/Ryutaro95/app/application/usecase"
	"github.com/Ryutaro95/app/infrastructure/persistence"
	"github.com/Ryutaro95/app/presentation/handler"
)

func InitializeHealthCheckHandler() *handler.HealthCheckHandler {
	healthCheckAccessor := persistence.NewHealthCheckPersistence()
	healthCheckUsecase := usecase.NewHealthCheckUsecase(healthCheckAccessor)
	healthCheckHandler := handler.NewHealthCheckHandler(healthCheckUsecase)

	return healthCheckHandler
}
