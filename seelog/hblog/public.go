package hblog

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

func Warnf(format string, params ...interface{}) error {
	g_logMutex.Lock()
	defer g_logMutex.Unlock()
	return g_myLogger.Warnf(format, params...)
}

func Errorf(format string, params ...interface{}) error {
	g_logMutex.Lock()
	defer g_logMutex.Unlock()
	return g_myLogger.Errorf(format, params...)
}

func Criticalf(format string, params ...interface{}) error {
	g_logMutex.Lock()
	defer g_logMutex.Unlock()
	return g_myLogger.Criticalf(format, params...)
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

func Warn(v ...interface{}) error {
	g_logMutex.Lock()
	defer g_logMutex.Unlock()
	return g_myLogger.Warn(v...)
}

func Error(v ...interface{}) error {
	g_logMutex.Lock()
	defer g_logMutex.Unlock()
	return g_myLogger.Error(v...)
}

func Critical(v ...interface{}) error {
	g_logMutex.Lock()
	defer g_logMutex.Unlock()
	return g_myLogger.Critical(v...)
}

func Flush() {
	g_logMutex.Lock()
	defer g_logMutex.Unlock()
	g_myLogger.Flush()
}
