package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func json_main() {
	fmt.Println("start")
	r := &json_Reader{d: `{"A":"b","C":10}`, i: 0}
	d := json.NewDecoder(r)
	t := &json_Data{}
	d.Decode(t)
	fmt.Println(*t)
}

type json_Reader struct {
	d string
	i int
}

func (r *json_Reader) Read(p []byte) (int, error) {
	time.Sleep(500 * time.Millisecond)
	p[0] = r.d[r.i]
	r.i++
	fmt.Println(p[0])
	return 1, nil
}

type json_Data struct {
	A string
	C int
}
