package platform

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Logger *logrus.Logger

func init() {

	Logger = &logrus.Logger{}

	Logger.Info("Failed to logs to file, using default stderr")

	file, err := os.OpenFile("./sdk.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		Logger.Out = file
	} else {
		Logger.Info("Failed to logs to file, using default stderr")
	}

	// 设置日志格式为JSON格式
	Logger.SetFormatter(&logrus.JSONFormatter{})

	// 设置日志级别为InfoLevel
	Logger.SetLevel(logrus.InfoLevel)
}
