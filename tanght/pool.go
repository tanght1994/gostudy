package main

import (
	"bytes"
	"fmt"
	"sync"
	"time"
)

var gTimes = 800 * 10000
var gPool = sync.Pool{New: func() interface{} {
	return bytes.NewBuffer(make([]byte, 2048))
}}

func pool_main() {
	t0 := time.Now().Unix()
	pool_test02()
	t1 := time.Now().Unix()
	fmt.Println(t1 - t0)
}

func pool_test01() {
	wg := sync.WaitGroup{}
	for i := 0; i < gTimes; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			b := bytes.NewBuffer(make([]byte, 2048))
			b.WriteString("1")
		}()
	}
}

func pool_test02() {
	wg := sync.WaitGroup{}
	for i := 0; i < gTimes; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			b := gPool.Get()
			c := b.(*bytes.Buffer)
			c.WriteString("1")
		}()
	}
}
