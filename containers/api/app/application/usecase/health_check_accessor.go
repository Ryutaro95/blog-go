package usecase

type HealthCheckAccessor interface {
	HealthCheckMysql() error
}
