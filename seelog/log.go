package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/cihub/seelog"
)

const G_LogCheckInterval = 100 * time.Millisecond

func replaceLogger(path string) error {
	logger, err := seelog.LoggerFromConfigAsFile(path)
	if err != nil {
		return fmt.Errorf("create logger error, %s", err.Error())
	}
	return seelog.ReplaceLogger(logger)
}

func filemd5(path string) (string, error) {
	bys, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	bmd5 := md5.Sum(bys)
	return hex.EncodeToString(bmd5[:]), nil
}

func initLoggerFromFile(path string) error {
	md5, err := filemd5(path)
	if err != nil {
		return err
	}
	err = replaceLogger(path)
	if err != nil {
		return err
	}
	go monitorLoggerConfig(path, G_LogCheckInterval, md5)
	return nil
}

func monitorLoggerConfig(path string, interval time.Duration, lastmd5 string) {
	for {
		md5, err := filemd5(path)
		if err == nil {
			if md5 != lastmd5 {
				lastmd5 = md5
				seelog.Critical("logger config changed")
				err = replaceLogger(path)
				if err != nil {
					seelog.Criticalf("logger update error, %s", err.Error())
				} else {
					seelog.Critical("logger update successful")
				}
			}
		} else {
			seelog.Criticalf("logger config md5 error, %s", err.Error())
		}
		seelog.Flush()
		time.Sleep(interval)
	}
}
