package main

import (
	"crypto/sha1"
	"fmt"
	"io"

	_ "github.com/go-sql-driver/mysql"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	s := "tanght niubi"
	h := sha1.New()
	io.WriteString(h, s)
	v1 := h.Sum(nil)
	v2 := sha1.Sum([]byte(s))
	fmt.Println(v1)
	fmt.Println(v2)
}
