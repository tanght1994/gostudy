package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"sync"
	"time"
)

var g_foods int
var g_cond sync.Cond

func main() {
	g_cond = *sync.NewCond(new(sync.Mutex))
	for i := 0; i < 3; i++ {
		go eatten()
	}
	go makefooduntilendtheworld()
	time.Sleep(time.Hour)
}

func eatone(id uint64) {
	g_cond.L.Lock()
	defer g_cond.L.Unlock()
	for g_foods == 0 {
		fmt.Println(id, " wait food")
		g_cond.Wait()
	}
	g_foods = g_foods - 1
	fmt.Println(id, " eat one")
}

func eatten() {
	for i := 0; i < 10; i++ {
		id := GetGoroutineID()
		eatone(id)
	}
}

func makefood() {
	g_cond.L.Lock()
	defer g_cond.L.Unlock()
	add := (rand.Int() % 10) + 1
	g_foods = g_foods + add
	fmt.Println("make ", add, " food")
	g_cond.Signal()
}

func makefooduntilendtheworld() {
	for {
		makefood()
		time.Sleep(1 * time.Second)
	}
}

func GetGoroutineID() uint64 {
	b := make([]byte, 64)
	runtime.Stack(b, false)
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
