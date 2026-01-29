package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// TODO: добавить запись логов в графану
// TODO: добавить разделение на микросервис
func Init() *logrus.Logger {
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	logger.SetOutput(os.Stdout)
	return logger
}
