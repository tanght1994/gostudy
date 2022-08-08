package main

import "fmt"

func main() {
	list := []int{3, 5, 1}
	fmt.Println(search(list, 5))
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	if nums[0] < nums[len(nums)-1] {
		return bsearch(nums, target)
	}
	l, r := 0, len(nums)-1
	sep := -1
	for l < r {
		if r-l == 1 {
			if nums[l] < nums[r] {
				sep = l
			} else {
				sep = r
			}
			break
		}
		m := l + (r-l)/2
		if nums[m] > nums[l] {
			l = m
		} else {
			r = m
		}
	}

	if target == nums[0] {
		return 0
	} else if target > nums[0] {
		return bsearch(nums[0:sep-1], target)
	} else {
		res := bsearch(nums[sep:], target)
		if res == -1 {
			return -1
		}
		return sep + res
	}
}

func bsearch(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}
	l, r := 0, len(nums)-1
	for l < r {
		if r == l+1 {
			if nums[r] == target {
				return r
			}
			if nums[l] == target {
				return l
			}
			return -1
		}
		m := l + (r-l)/2
		if nums[m] > target {
			r = m
		} else if nums[m] < target {
			l = m
		} else {
			return m
		}
	}
	return -1
}
