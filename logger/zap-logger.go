package logger

import (
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
	encoderCfg.EncodeTime = zapcore.RFC3339TimeEncoder

	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:       false,
		DisableCaller:     true,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          "json",
		EncoderConfig:     encoderCfg,
		OutputPaths: []string{
			"log-history.log",
		},
		ErrorOutputPaths: []string{
			"log-history.log",
		},
		InitialFields: map[string]any{},
	}

	logger := &ZapLogger{
		log: zap.Must(config.Build()),
	}

	return *logger
}

func (zlog ZapLogger) LogInfo(msg string, fields ...map[string]any) {
	logFields := []zapcore.Field{}
	for k, v := range fields[0] {
		logFields = append(logFields, zap.Any(k, v))
	}
	zlog.log.Info(msg, logFields...)
}

func (zlog ZapLogger) LogErr(msg string, fields ...map[string]any) {
	logFields := []zapcore.Field{}
	for k, v := range fields[0] {
		logFields = append(logFields, zap.Any(k, v))
	}
	zlog.log.Error(msg, logFields...)
}

func (zlog ZapLogger) LogDebug(msg string, fields ...map[string]any) {
	logFields := []zapcore.Field{}
	for k, v := range fields[0] {
		logFields = append(logFields, zap.Any(k, v))
	}
	zlog.log.Debug(msg, logFields...)
}

func (zlog ZapLogger) LogWarn(msg string, fields ...map[string]any) {
	logFields := []zapcore.Field{}
	for k, v := range fields[0] {
		logFields = append(logFields, zap.Any(k, v))
	}
	zlog.log.Warn(msg, logFields...)
}

func (zlog ZapLogger) LogFatal(msg string, fields ...map[string]any) {
	logFields := []zapcore.Field{}
	for k, v := range fields[0] {
		logFields = append(logFields, zap.Any(k, v))
	}
	zlog.log.Fatal(msg, logFields...)
}
