package main

import (
	"fmt"
	"sort"
)

func sort_main() {
	a := []int{1, 8, 4, 5, 6}
	sort.Sort(sort.IntSlice(a))
	fmt.Println(a)
}
