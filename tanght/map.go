package main

import (
	"io"
	"os"
	"time"
)

type map_tanght struct {
	i uint
	s string
}

func (m *map_tanght) Read(p []byte) (n int, err error) {
	p[0] = m.s[m.i]
	m.i = (m.i + 1) % uint(len(m.s))
	time.Sleep(time.Second)
	return 1, nil
}

func map_main() {
	m := &map_tanght{s: "abcdefghijklmnopqrstuvwxyz"}
	io.Copy(os.Stdout, m)
}
