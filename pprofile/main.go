package main

import (
	"fmt"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go run()
	http.ListenAndServe("0.0.0.0:8000", nil)
}

func run() {
	for {
		dur := time.Duration(rand.Intn(1000)) * time.Millisecond
		time.Sleep(dur)
		go worker(dur)
	}
}

func worker(dur time.Duration) {
	a := make([]int, rand.Intn(1023)+1)
	fmt.Println(len(a))
	time.Sleep(dur)
}
