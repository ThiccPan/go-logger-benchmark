package logger

const zapLogfilePath = "zap-log-history.log"
const logrusLogfilePath = "logrus-log-history.log"

type Ilogger interface {
	LogInfo(msg string, fields ...map[string]any)
	LogErr(msg string, fields ...map[string]any)
	LogDebug(msg string, fields ...map[string]any)
	LogWarn(msg string, fields ...map[string]any)
	LogFatal(msg string, fields ...map[string]any)
}