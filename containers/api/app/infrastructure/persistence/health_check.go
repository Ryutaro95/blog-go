package persistence

import (
	"github.com/Ryutaro95/app/infrastructure/logger"
	"go.uber.org/zap"
)

type HealthCheckPersistence struct {
	Text string
}

func NewHealthCheckPersistence() *HealthCheckPersistence {
	return &HealthCheckPersistence{
		Text: "hoge",
	}
}

func (hc HealthCheckPersistence) HealthCheckMysql() error {
	logger.Debug("sql query", zap.String("query", "test"))

	return nil
}
