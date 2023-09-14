package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Logger *logrus.Logger

func init() {

	Logger = &logrus.Logger{}

	Logger.Info("Failed to logs to file, using default stderr")

	file, err := os.OpenFile("/tmp/logs/platform-sdk.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		Logger.Out = file
	} else {
		Logger.Info("Failed to logs to file, using default stderr")
	}

	Logger.SetFormatter(&logrus.JSONFormatter{})

	Logger.SetLevel(logrus.InfoLevel)
}
