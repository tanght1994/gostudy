package main

import (
	"log"

	"github.com/cihub/seelog"
)

func main() {
	logger, err := seelog.LoggerFromConfigAsFile("seelog.xml")
	if err != nil {
		log.Fatal(1)
	}
	logger.Info("abc")
	logger.Info("1")
	logger.Info("2")
	logger.Info("3")
	logger.Close()
}
