package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/cihub/seelog"
)

const G_LogConfigPath = "seelog.xml"

func main() {
	logger, err := seelog.LoggerFromConfigAsFile("seelog.xml")
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 1000; i++ {
		logger.Trace("Trace")
		logger.Debug("Debug")
		logger.Info("Info")
		logger.Warn("Warn")
		logger.Error("Error")
		logger.Flush()
		fmt.Println("-------------------")
		time.Sleep(1 * time.Second)
	}
}

func replaceLogger() error {
	logger, err := seelog.LoggerFromConfigAsFile(G_LogConfigPath)
	if err != nil {
		return fmt.Errorf("create logger error, %s", err.Error())
	}
	return seelog.ReplaceLogger(logger)
}

type monitorLogConfig struct {
	lastmd5 string
	path    string
}

func (m *monitorLogConfig) monitor() {
	bys, err := ioutil.ReadFile(m.path)
	if err != nil {
		// log
		return
	}
	t1 := md5.Sum(bys)
	t2 := hex.EncodeToString(t1[:])
	if t2 == m.lastmd5 {
		return
	}
	m.lastmd5 = t2
	replaceLogger()
}
