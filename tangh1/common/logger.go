package common

import (
	"fmt"
	"log"
	"os"
	"sync/atomic"
)

const (
	Ldebug = iota
	Linfo
	Lerror
)

type logger struct {
	l     *log.Logger
	level int32
}

func (m *logger) SetLevel(level int32) {
	atomic.StoreInt32(&m.level, level)
}

func (m *logger) log(level int32, format string, v ...interface{}) {
	if level < atomic.LoadInt32(&m.level) {
		return
	}
	lable := "   NO"
	if level == Ldebug {
		lable = "DEBUG"
	} else if level == Linfo {
		lable = " INFO"
	} else if level == Lerror {
		lable = "ERROR"
	}
	format = fmt.Sprintf("%s %s", lable, format)
	m.l.Printf(format, v...)
}

func (m *logger) Info(msg string) {
	m.log(Linfo, "%s\n", msg)
}

func (m *logger) Infof(format string, v ...interface{}) {
	m.log(Linfo, format, v...)
}

func newLogger() (*logger, error) {
	f, e := os.OpenFile("log.log", os.O_APPEND, 0666)
	if e != nil {
		f, e = os.OpenFile("log.log", os.O_CREATE, 0666)
		if e != nil {
			return nil, e
		}
	}
	return &logger{log.New(f, "", log.LstdFlags), Ldebug}, nil
}
