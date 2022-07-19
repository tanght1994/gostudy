package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(searchRange([]int{2, 2}, 3))
}

func must(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

}
