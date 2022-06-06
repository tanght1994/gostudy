package main

import (
	"fmt"
)

func main() {
	if true {
		fmt.Println(1)
		defer func() {
			fmt.Println(123)
		}()
		fmt.Println(2)
	}
	fmt.Println(3)
}
