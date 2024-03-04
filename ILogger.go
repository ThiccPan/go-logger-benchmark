package main

type Ilogger interface {
	logInfo(msg string)
	logErr(msg string)
}