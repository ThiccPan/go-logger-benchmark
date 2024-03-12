package logger

import (
	"os"

	"github.com/rs/zerolog"
)

type ZerologLogger struct {
	log *zerolog.Logger
}

func InitZerolog() ZerologLogger {
	// setup logging file location
	logFile, err := os.OpenFile("./log-history.log", os.O_RDWR, 0644)
	if err != nil {
		panic(1)
	}

	// init new logger instance with the called output fields
	logger := zerolog.New(logFile)
	logger = logger.With().Timestamp().Logger()
	// logger = logger.With().Caller().Logger()

	logger.Debug().Msg("logger created")

	return ZerologLogger{
		log: &logger,
	}
}

func (zlog ZerologLogger) LogInfo(msg string) {
	zlog.log.Info().Msg(msg)
}

func (zlog ZerologLogger) LogErr(msg string) {
	zlog.log.Error().Msg(msg)
}
