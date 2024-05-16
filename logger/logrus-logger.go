package logger

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

func InitLogrusLogger() *logrus.Logger {
	// open logfile location
	logFile, err := os.OpenFile(logrusLogfilePath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic(1)
	}

	// setup log formatter
	formatter := new(logrus.JSONFormatter)
	formatter.TimestampFormat = time.RFC3339

	formatter.FieldMap = logrus.FieldMap{
		logrus.FieldKeyMsg: "message",
	}

	logger := &logrus.Logger{
		Out:       logFile,
		Formatter: formatter,
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.DebugLevel,
	}
	return logger
}
