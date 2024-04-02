package logger

type Ilogger interface {
	LogInfo(msg string, fields ...map[string]any)
	LogErr(msg string, fields ...map[string]any)
	LogDebug(msg string, fields ...map[string]any)
	LogWarn(msg string, fields ...map[string]any)
	LogFatal(msg string, fields ...map[string]any)
}