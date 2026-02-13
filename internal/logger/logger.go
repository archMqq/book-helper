package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// TODO: добавить запись логов в графану
func InitForService(serviceName string) *logrus.Entry {
	return Init().WithField("app-service", serviceName)
}

func Init() *logrus.Logger {
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	logger.SetOutput(os.Stdout)
	return logger
}
