package main

import (
	"crypto/sha1"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	h := sha1.New()
	h.Sum(nil)

	res := fun1(nil)
	fmt.Println(res)
}

func fun1(aaa []int) []int {
	haha := append(aaa, 1, 2)
	return haha
}
