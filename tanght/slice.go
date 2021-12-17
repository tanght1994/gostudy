package main

import "fmt"

func slice_main() {
	a := make([]int, 5, 10)
	fmt.Println(a)
	b := a[:20]
	fmt.Println(b)
	_ = append(a, 1, 2, 3)
	fmt.Println(b)
}
