package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var wg = &sync.WaitGroup{}
var a int32 = 0

func main() {
	t0 := time.Now().UnixMilli()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go aaa()
	}
	wg.Wait()
	t1 := time.Now().UnixMilli()
	t := (t1 - t0)
	fmt.Println(a, t)
}

func aaa() {
	defer wg.Done()
	for i := 0; i < 1000000; i++ {
		atomic.LoadInt32(&a)
	}
}
