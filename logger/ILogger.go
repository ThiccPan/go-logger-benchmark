package logger

type Ilogger interface {
	LogInfo(msg string)
	LogErr(msg string)
	LogDebug(msg string)
	LogWarn(msg string)
	LogFatal(msg string)
}