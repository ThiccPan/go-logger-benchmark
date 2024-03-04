package main

import "go.uber.org/zap"

type ZapLogger struct {
	log *zap.Logger
}

// logErr implements Ilogger.

func InitZap() ZapLogger {
	logger := zap.Must(zap.NewProduction())
	// defer logger.Sync()
	a := ZapLogger{log: logger}
	logger.Info("logger init successfull")
	return a
}

func (zlog ZapLogger) logInfo(msg string) {
	zlog.log.Info(msg)
}

func (zlog ZapLogger) logErr(msg string) {
	zlog.log.Error(msg)
}
