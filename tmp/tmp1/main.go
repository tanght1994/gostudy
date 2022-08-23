package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	fmt.Println(maxSlidingWindow([]int{7, 2, 4}, 2))
}

func maxSlidingWindow(nums []int, k int) []int {
	res := []int{}
	que := []int{}
	que = append(que, nums[0])
	for i := 1; i < k; i++ {
		for len(que) != 0 {
			if que[len(que)-1] < nums[i] {
				que = que[:len(que)-1]
			} else {
				break
			}
		}
		que = append(que, nums[i])
	}
	res = append(res, que[0])
	for i := k; i < len(nums); i++ {
		if que[len(que)-1] == nums[i-k] {
			que = que[1:]
		}
		for len(que) != 0 {
			if que[len(que)-1] < nums[i] {
				que = que[:len(que)-1]
			} else {
				break
			}
		}
		que = append(que, nums[i])
		res = append(res, que[0])
	}
	return res
}
