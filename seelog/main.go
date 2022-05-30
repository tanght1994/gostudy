package main

import (
	"tanght/hblog"
	"time"
)

func main() {
	work()
	for {
		time.Sleep(1 * time.Second)
	}
}

func work() {
	for i := 0; i < 100; i++ {
		hblog.Trace("Trace")
		hblog.Debug("Debug")
		hblog.Info("Info")
		hblog.Warn("Warn")
		hblog.Error("Error")
		hblog.Critical("-------------------")
		time.Sleep(1 * time.Second)
	}
	hblog.Flush()
}
