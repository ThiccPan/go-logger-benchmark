package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

type LogrusLogger struct {
	log *logrus.Logger
}

func InitLogrusLogger() *LogrusLogger {
	logger := &LogrusLogger{
		log: &logrus.Logger{
			Out:       os.Stderr,
			Formatter: new(logrus.TextFormatter),
			Hooks:     make(logrus.LevelHooks),
			Level:     logrus.DebugLevel,
		},
	}
	return logger
}

func (llog *LogrusLogger) LogInfo(msg string) {
	llog.log.Info(msg)
}

func (llog *LogrusLogger) LogErr(msg string) {
	llog.log.Error(msg)
}

func (llog *LogrusLogger) LogDebug(msg string) {
	llog.log.Debug(msg)
}

func (llog *LogrusLogger) LogWarn(msg string) {
	llog.log.Warn(msg)
}

func (llog *LogrusLogger) LogFatal(msg string) {
	llog.log.Fatal(msg)
}