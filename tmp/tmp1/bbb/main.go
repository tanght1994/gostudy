package main

import (
	"fmt"
)

func main() {
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
}

type Info struct {
	l, r int
}
type Item struct {
	h, i int
}

func largestRectangleArea(heights []int) int {
	infos := make([]Info, len(heights))
	stack := make([]Item, 0)
	for i := 0; i < len(heights); i++ {
		if len(stack) == 0 {
			stack = append(stack, Item{heights[i], i})
			continue
		}
		for len(stack) != 0 && heights[i] < stack[len(stack)-1].h {
			infos[stack[len(stack)-1].i].r = i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, Item{heights[i], i})
	}
	for len(stack) != 0 {
		infos[stack[len(stack)-1].i].r = len(heights)
		stack = stack[:len(stack)-1]
	}

	for i := len(heights) - 1; i >= 0; i-- {
		if len(stack) == 0 {
			stack = append(stack, Item{heights[i], i})
			continue
		}
		for len(stack) != 0 && heights[i] < stack[len(stack)-1].h {
			infos[stack[len(stack)-1].i].l = i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, Item{heights[i], i})
	}
	for len(stack) != 0 {
		infos[stack[len(stack)-1].i].l = -1
		stack = stack[:len(stack)-1]
	}

	max := 0
	for i := range infos {
		area := heights[i] * (infos[i].r - infos[i].r - 1)
		if area > max {
			max = area
		}
	}

	return max
}
