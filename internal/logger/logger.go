package logger

import (
	"time"

	"github.com/sirupsen/logrus"
	glog "gorm.io/gorm/logger"
)

var logger = logrus.New()

func InitLogger() {
	logger.SetLevel(logrus.TraceLevel)
	return
}

func Logger() *logrus.Logger {
	return logger
}

func NewGormLogger() glog.Interface {
	gormCfg := glog.Config{
		SlowThreshold: time.Second,
		Colorful:      true,
		LogLevel:      glog.Info,
	}
	return glog.New(
		logger,
		gormCfg,
	)
}
