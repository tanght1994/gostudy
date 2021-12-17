package main

import "fmt"

func byte_main() {
	var a uint16 = 65280
	fmt.Println(byte(a))
	fmt.Println(byte(a >> 8))
}
