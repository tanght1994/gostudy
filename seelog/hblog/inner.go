package hblog

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
	"time"

	"github.com/cihub/seelog"
)

var g_myLogger seelog.LoggerInterface // 全局logger，非并发安全，必须加普通锁，不能是读写锁
var g_logMutex *sync.Mutex            // 操作(读和写两种操作)g_myLogger的时候，必须持有这个锁
// var g_onceinit sync.Once              // InitLog只能调用一次，放弃Once，由用户负责只调用一次

func init() {
	customizeSeelog()
	g_logMutex = new(sync.Mutex)
	g_myLogger = seelog.Disabled
}

func InitLog(config_path string, monitor_interval time.Duration) error {
	if monitor_interval < time.Second {
		monitor_interval = 1 * time.Second
	}
	err := replaceLoggerFromFile(config_path)
	if err != nil {
		return err
	}
	md5, err := filemd5(config_path)
	if err != nil {
		return err
	}
	go monitorConfig(config_path, monitor_interval, md5)
	return nil
}

func replaceLoggerFromFile(path string) error {
	errmsg := ""
	logger, err := seelog.LoggerFromConfigAsFile(path)
	if err != nil {
		errmsg = fmt.Sprintf("create logger from config file failed, %s", err.Error())
		g_myLogger.Critical(errmsg)
		tostderr(errmsg)
		return err
	}
	// 告诉老哥调用栈帧增加了一层，让老哥寻找栈帧的时候不要出错
	logger.SetAdditionalStackDepth(1)
	err = replaceLogger(logger)
	if err != nil {
		errmsg = fmt.Sprintf("replace hblog default logger failed, %s", err.Error())
		g_myLogger.Critical(errmsg)
		tostderr(errmsg)
	} else {
		errmsg = "replace hblog default logger sucessful!"
		g_myLogger.Critical(errmsg)
		tostderr(errmsg)
	}
	return err
}

func replaceLogger(logger seelog.LoggerInterface) error {
	if logger == nil {
		return errors.New("logger can not be nil")
	}
	g_logMutex.Lock()
	defer g_logMutex.Unlock()
	defer func() {
		if err := recover(); err != nil {
			fmt.Fprintf(os.Stderr, "recovered from panic during ReplaceLogger: %s", err)
		}
	}()

	if g_myLogger != nil && !g_myLogger.Closed() && g_myLogger != seelog.Disabled {
		g_myLogger.Flush()
		g_myLogger.Close()
	}

	g_myLogger = logger

	return nil
}

func monitorConfig(path string, interval time.Duration, lastmd5 string) {
	for {
		md5, err := filemd5(path)
		if err == nil {
			if md5 != lastmd5 {
				lastmd5 = md5
				g_myLogger.Critical("log config changed")
				replaceLoggerFromFile(path)
			}
		} else {
			g_myLogger.Criticalf("calculate log config md5 error, %s", err.Error())
		}
		time.Sleep(interval)
	}
}

func filemd5(path string) (string, error) {
	bys, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	bmd5 := md5.Sum(bys)
	return hex.EncodeToString(bmd5[:]), nil
}

func customizeSeelog() {
	seelog.RegisterCustomFormatter("ThtDateTime", func(param string) seelog.FormatterFunc {
		return func(message string, level seelog.LogLevel, context seelog.LogContextInterface) interface{} {
			return context.CallTime().Format("2006-01-02 15:04:05 -07:00")
		}
	})
}

func tostderr(msg string) {
	fmt.Fprint(os.Stderr, msg+"\n")
}
