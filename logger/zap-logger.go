package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// logErr implements Ilogger.
func InitZap() *zap.Logger {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "time"
	encoderCfg.MessageKey = "message"
	encoderCfg.EncodeTime = zapcore.RFC3339TimeEncoder

	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:       false,
		DisableCaller:     true,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          "json",
		EncoderConfig:     encoderCfg,
		OutputPaths: []string{
			zapLogfilePath,
		},
		ErrorOutputPaths: []string{
			zapLogfilePath,
		},
		InitialFields: map[string]any{},
	}

	logger := zap.Must(config.Build())

	return logger
}
