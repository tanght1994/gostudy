package main

import (
	"fmt"
)

func defer_main() {
	a := defer_test01()
	fmt.Println(a)
}

func defer_test01() (i int) {
	defer func() {
		i++
	}()
	return 1
}
