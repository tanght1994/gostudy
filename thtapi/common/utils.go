package common

import (
	"fmt"
	"sync/atomic"
	"time"
)

const (
	LevelDebug    = int32(0)
	LevelInfo     = int32(1)
	LevelWarn     = int32(2)
	LevelError    = int32(3)
	LevelCritical = int32(4)
)

var logLevel atomic.Int32

func log(level int32, msg string) {
	if level < logLevel.Load() {
		return
	}
	logmsg := fmt.Sprintf("%s %s\n", time.Now().Format("2006-01-02 15:04:05 -0700"), msg)
	fmt.Print(logmsg)
}

func LogDebug(msg string) {
	log(LevelDebug, "[D] "+msg)
}

func LogInfo(msg string) {
	log(LevelInfo, "[I] "+msg)
}

func LogWarn(msg string) {
	log(LevelWarn, "[W] "+msg)
}

func LogError(msg string) {
	log(LevelError, "[E] "+msg)
}

func LogCritical(msg string) {
	log(LevelCritical, "[C] "+msg)
}

func SetLogLevel(level int32) {
	logLevel.Store(level)
}
