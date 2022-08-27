package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.WarnLevel)
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info(`A walrus "haha" appears`)
	log.Error("123", "haha", "hehehe")
}
