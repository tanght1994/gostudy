package main

import (
	"fmt"
	"unsafe"
)

func string_main() {
	b := []byte{104, 101, 108, 108, 111}
	p := unsafe.Pointer(&b)
	s := *(*string)(p)
	fmt.Println(&b, &s)
	fmt.Printf("%d %d", &b, &s)
	println(&b, &s)
}
