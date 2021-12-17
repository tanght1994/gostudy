package main

import (
	"bytes"
	"fmt"
)

func ttt_main() {
	s1 := []int{1, 2, 3, 4}
	s2 := []int{7, 8}
	fmt.Println(s2)
	s2 = s1
	fmt.Println("--------")
	fmt.Println(s1, s2)
	s2[0] = 100
	fmt.Println(s1, s2)
}

func ttt_t1() {
	s1 := "tanghthaha tanght 666"
	s2 := bytes.Replace([]byte(s1), []byte("tanght"), []byte("www"), -1)
	fmt.Println(s1)
	fmt.Println(string(s2))
}
