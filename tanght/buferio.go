package main

import (
	"bufio"
	"fmt"
	"time"
)

func bufferio_main() {
	bufio_test01()
}

type Product struct {
	idx  int
	item []byte
}

func (s *Product) Read(p []byte) (n int, err error) {
	time.Sleep(2 * time.Second)
	p[0] = s.item[s.idx]
	s.idx = (s.idx + 1) % len(s.item)
	return 1, nil
}

func bufio_test01() {
	p := &Product{idx: 0, item: []byte("abcdef\nghijk")}
	rd := bufio.NewReader(p)
	l, err := rd.ReadSlice('\n')
	fmt.Println(l, err)
}
