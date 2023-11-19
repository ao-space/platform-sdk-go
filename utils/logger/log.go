package logger

import (
	"go.uber.org/zap"
)

func Default() *zap.SugaredLogger {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	return logger.Sugar()
}
