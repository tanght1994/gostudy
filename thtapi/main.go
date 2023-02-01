package main

import (
	"fmt"
	"net"
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
	ip, ipnet, e := net.ParseCIDR("192.168.0.0/16")
	fmt.Println(ip)
	fmt.Println(ipnet)
	fmt.Println(e)
	fmt.Println(ipnet.Contains(net.ParseIP("192.168.7.45")))
	fmt.Println(ipnet.Contains(net.ParseIP("192.168.8.45")))
	fmt.Println(ipnet.Contains(net.ParseIP("192.169.7.45")))
	fmt.Println(ipnet.Contains(net.ParseIP("192.168.7.47")))
	return
	common.SetLogLevel(common.LevelDebug)
	common.LogCritical("thtapi start")
	transport := http.Transport{
		IdleConnTimeout: 10 * time.Second,
	}
	Requester = http.Client{
		Transport: &transport,
	}
	gin.SetMode(gin.ReleaseMode)
	GinEngine = gin.New()
	registerHandleFunc()

	if err := GinEngine.Run("0.0.0.0:8000"); err != nil {
		common.LogError(`GinEngine.Run error, ` + err.Error())
	}
}

func must(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func registerHandleFunc() {
	GinEngine.Any("/", proxy)
	GinEngine.POST("/interval/set_endpoint", setEndPoint)
}
