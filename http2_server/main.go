package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		now := time.Now().Unix()
		fmt.Println(now, req.RemoteAddr)
		time.Sleep(time.Duration(rand.Intn(10)+8) * time.Second)
		resp.Write([]byte(fmt.Sprintf("%d hello world", now)))
	})
	e := http.ListenAndServeTLS(":8000", "server.crt", "server.key", mux)
	if e != nil {
		fmt.Println(e)
	}
}
