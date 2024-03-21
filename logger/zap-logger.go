package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
	log *zap.Logger
}

// logErr implements Ilogger.
func InitZap() ZapLogger {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "time"
	encoderCfg.MessageKey = "message"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:       false,
		DisableCaller:     true,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          "json",
		EncoderConfig:     encoderCfg,
		OutputPaths: []string{
			"stderr",
			"log-history.log",
		},
		ErrorOutputPaths: []string{
			"stderr",
			"log-history.log",
		},
		InitialFields: map[string]interface{}{
			"pid": os.Getpid(),
		},
	}

	logger := &ZapLogger{
		log: zap.Must(config.Build()),
	}

	return *logger
}

func (zlog ZapLogger) LogInfo(msg string) {
	zlog.log.Info(msg)
}

func (zlog ZapLogger) LogErr(msg string) {
	zlog.log.Error(msg)
}

func (zlog ZapLogger) LogDebug(msg string) {
	zlog.log.Debug(msg)
}

func (zlog ZapLogger) LogWarn(msg string) {
	zlog.log.Warn(msg)
}

func (zlog ZapLogger) LogFatal(msg string) {
	zlog.log.Fatal(msg)
}