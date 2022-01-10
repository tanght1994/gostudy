package main

import (
	"log"
	"time"

	"github.com/cihub/seelog"
)

const G_LogConfigPath = "seelog.xml"

func main() {
	if err := initLoggerFromFile(G_LogConfigPath); err != nil {
		log.Fatal("init log error, ", err.Error())
	}

	for i := 0; i < 3; i++ {
		go work()
	}

	for {
		time.Sleep(1 * time.Second)
	}
}

func work() {
	for {
		seelog.Trace("Trace")
		seelog.Debug("Debug")
		seelog.Info("Info")
		seelog.Warn("Warn")
		seelog.Error("Error")
		seelog.Critical("-------------------")
		time.Sleep(1 * time.Second)
	}
}
