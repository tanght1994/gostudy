package main

import (
	"bufio"
	"bytes"
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
	bufio.NewReader()
	bytes.NewReader()
}
