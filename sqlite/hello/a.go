package hello

import (
	"fmt"
	"time"
)

var g *int

func init() {
	t := 8
	g = &t
	go print_g()
}

func print_g() {
	for {
		fmt.Println(g == nil)
		time.Sleep(1 * time.Second)
	}
}
