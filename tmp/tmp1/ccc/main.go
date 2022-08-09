package main

import (
	"fmt"
	"os"
)

func main() {
	a := map[string]int{}
	hahahahahahaha(a)
	fmt.Println(a)
}

func hahahahahahaha(data map[string]int) {
	data["a"] = 1
	data["b"] = 1
	data["c"] = 1
	data["d"] = 1
	data["e"] = 1
	data["f"] = 1
	data["g"] = 1
}

func must(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
