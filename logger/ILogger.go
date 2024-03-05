package logger

type Ilogger interface {
	LogInfo(msg string)
	LogErr(msg string)
}