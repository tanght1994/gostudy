package main

import (
	"fmt"
	"time"
)

func main() {
	time.AfterFunc(5*time.Second, func() {
		fmt.Println("1")
	})
	time.Sleep(time.Minute)
}
