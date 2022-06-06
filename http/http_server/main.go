package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/tanght/", func(resp http.ResponseWriter, req *http.Request) {
		// aaa := []byte{}
		// a := aaa[100]
		// fmt.Println(a)
		now := time.Now().Unix()
		fmt.Println(now, req.RemoteAddr)
		time.Sleep(time.Duration(rand.Intn(10)+8) * time.Second)
		resp.Write([]byte(fmt.Sprintf("%d hello world", now)))
	})
	http.ListenAndServe("0.0.0.0:8000", mux)
}
