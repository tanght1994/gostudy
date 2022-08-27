package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}

func aaa() {
	errtype := []string{"ERROR", "INFO", "INFO", "INFO", "INFO", "INFO", "WARING", "WARING", "DEBUG", "DEBUG", "CRIT"}
	gameid := []string{"1", "4003", "4003", "4003", "4003", "1009", "7788", "7788", "7788", "4004", "40053", "5566", "10065", "5566", "4000", "4001"}
	userid := []string{"100001", "100002", "100003", "100004", "100005", "100006", "100007", "100008", "100009", "100001", "100001", "100001", "100001", "100002"}
	log := []string{"1", "2", "3", "4", "5"}
	logtype := []string{"t1", "t2", "t3", "t4", "t5", "t6", "t7", "t8", "t9", "t1", "t1", "t1", "t2", "t2", "t3"}
}

func createdata(start, end time.Time, count int) {
	time.Now()
}