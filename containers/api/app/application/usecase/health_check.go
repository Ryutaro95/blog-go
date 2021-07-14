package usecase

import "github.com/Ryutaro95/app/infrastructure/logger"

type HealthCheckUsecaseAccessor interface {
	HealthCheck() error
}

type HealthCheckUsecase struct {
	healthCheckAccessor HealthCheckAccessor
}

func NewHealthCheckUsecase(hc HealthCheckAccessor) *HealthCheckUsecase {
	return &HealthCheckUsecase{
		healthCheckAccessor: hc,
	}
}

func (hc HealthCheckUsecase) HealthCheck() error {
	logger.Info("start mysql healtch check")

	if err := hc.healthCheckAccessor.HealthCheckMysql(); err != nil {
		return err
	}

	return nil
}
