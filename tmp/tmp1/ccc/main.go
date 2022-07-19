package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fs, err := os.Open("C:/Users/tanght/Desktop/tmp/gostudy/tmp/tmp1/tmp.txt")
	must(err)
	bys := make([]byte, 10)
	for i := 0; i < 1000000; i++ {
		n, err := fs.Read(bys)
		must(err)
		fmt.Println(n, " | ", string(bys))
		time.Sleep(100 * time.Millisecond)
	}
}

func must(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
