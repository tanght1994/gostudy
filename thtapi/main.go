package main

import (
	"fmt"
	"net/http"
	"os"
	"thtapi/common"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	// Requester ...
	Requester http.Client
	// GinEngine ...
	GinEngine *gin.Engine
)

func main() {
	common.SetLogLevel(common.LevelDebug)
	common.LogCritical("thtapi start")
	transport := http.Transport{
		IdleConnTimeout: 10 * time.Second,
	}
	Requester = http.Client{
		Transport: &transport,
	}
	registerHandleFunc()
	if err := http.ListenAndServe("0.0.0.0:8000", nil); err != nil {
		common.LogError(`http.ListenAndServe error, ` + err.Error())
	}
}

func must(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func registerHandleFunc() {
	GinEngine.Any("", nil)
	GinEngine.POST("/interval/set_endpoint", setEndPoint)
	http.HandleFunc("/", proxy)
	http.HandleFunc("/interval/set_svcaddr", nil)
}
