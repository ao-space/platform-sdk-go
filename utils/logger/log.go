package logger

import (
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"path/filepath"
)

var Logger *logrus.Logger

func init() {

	Logger = &logrus.Logger{}

	Logger.Info("Failed to logs to file, using default stderr")

	defaultPath := "/tmp/logs/platform-sdk.log"

	// 获取文件路径中的目录部分
	dir := filepath.Dir(defaultPath)

	// 检查目录是否存在，如果不存在则递归创建目录
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755) // 0755 是一个常用的权限设置，允许所有用户读取和执行，但只有所有者可以写入
		if err != nil {
			log.Fatalf("Error creating directory: %v", err)
		}
	}

	file, err := os.OpenFile("/tmp/logs/platform-sdk.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		Logger.Out = file
	} else {
		Logger.Info("Failed to logs to file, using default stderr")
	}

	Logger.SetFormatter(&logrus.JSONFormatter{})

	Logger.SetLevel(logrus.InfoLevel)
}
