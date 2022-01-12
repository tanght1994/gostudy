package hblog

import "time"

const (
	g_logConfigPath      = "seelog.xml"    // 配置文件路径
	g_logMonitorInterval = 1 * time.Second // 监控配置文件变换的时间间隔
)

func Tracef(format string, params ...interface{}) {
	g_logMutex.Lock()
	defer g_logMutex.Unlock()
	g_myLogger.Tracef(format, params...)
}

func Debugf(format string, params ...interface{}) {
	g_logMutex.Lock()
	defer g_logMutex.Unlock()
	g_myLogger.Debugf(format, params...)
}

func Infof(format string, params ...interface{}) {
	g_logMutex.Lock()
	defer g_logMutex.Unlock()
	g_myLogger.Infof(format, params...)
}

func Warnf(format string, params ...interface{}) {
	g_logMutex.Lock()
	defer g_logMutex.Unlock()
	g_myLogger.Warnf(format, params...)
}

func Errorf(format string, params ...interface{}) {
	g_logMutex.Lock()
	defer g_logMutex.Unlock()
	g_myLogger.Errorf(format, params...)
}

func Criticalf(format string, params ...interface{}) {
	g_logMutex.Lock()
	defer g_logMutex.Unlock()
	g_myLogger.Criticalf(format, params...)
	g_myLogger.Flush()
}

func Trace(v ...interface{}) {
	g_logMutex.Lock()
	defer g_logMutex.Unlock()
	g_myLogger.Trace(v...)
}

func Debug(v ...interface{}) {
	g_logMutex.Lock()
	defer g_logMutex.Unlock()
	g_myLogger.Debug(v...)
}

func Info(v ...interface{}) {
	g_logMutex.Lock()
	defer g_logMutex.Unlock()
	g_myLogger.Info(v...)
}

func Warn(v ...interface{}) {
	g_logMutex.Lock()
	defer g_logMutex.Unlock()
	g_myLogger.Warn(v...)
}

func Error(v ...interface{}) {
	g_logMutex.Lock()
	defer g_logMutex.Unlock()
	g_myLogger.Error(v...)
}

func Critical(v ...interface{}) {
	g_logMutex.Lock()
	defer g_logMutex.Unlock()
	g_myLogger.Critical(v...)
	g_myLogger.Flush()
}

func Flush() {
	g_logMutex.Lock()
	defer g_logMutex.Unlock()
	g_myLogger.Flush()
}
