package logger

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

type LogrusLogger struct {
	log *logrus.Logger
}

func InitLogrusLogger() *LogrusLogger {
	// open logfile location
	logFile, err := os.OpenFile("./log-history.log", os.O_RDWR, 0644)
	if err != nil {
		panic(1)
	}

	// setup log formatter
	formatter := new(logrus.JSONFormatter)
	formatter.TimestampFormat = time.RFC3339

	formatter.FieldMap = logrus.FieldMap{
		logrus.FieldKeyMsg: "message",
	}

	logger := &LogrusLogger{
		log: &logrus.Logger{
			Out:       logFile,
			Formatter: formatter,
			Hooks:     make(logrus.LevelHooks),
			Level:     logrus.DebugLevel,
		},
	}
	return logger
}

func (llog *LogrusLogger) LogInfo(msg string, fields ...map[string]any) {
	logFields := logrus.Fields{}
	for k, v := range fields[0] {
		logFields[k] = v
	}
	llog.log.WithFields(logFields).Info(msg)
}

func (llog *LogrusLogger) LogErr(msg string, fields ...map[string]any) {
	logFields := logrus.Fields{}
	for k, v := range fields[0] {
		logFields[k] = v
	}
	llog.log.WithFields(logFields).Error(msg)
}

func (llog *LogrusLogger) LogDebug(msg string, fields ...map[string]any) {
	logFields := logrus.Fields{}
	for k, v := range fields[0] {
		logFields[k] = v
	}
	llog.log.WithFields(logFields).Debug(msg)
}

func (llog *LogrusLogger) LogWarn(msg string, fields ...map[string]any) {
	logFields := logrus.Fields{}
	for k, v := range fields[0] {
		logFields[k] = v
	}
	llog.log.WithFields(logFields).Warn(msg)
}

func (llog *LogrusLogger) LogFatal(msg string, fields ...map[string]any) {
	logFields := logrus.Fields{}
	for k, v := range fields[0] {
		logFields[k] = v
	}
	llog.log.WithFields(logFields).Fatal(msg)
}
