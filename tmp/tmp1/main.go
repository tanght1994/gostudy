package main

import (
	"fmt"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	a := "aaaa"
	fmt.Println(longestPalindrome(a))
}

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	if len(s) == 2 {
		if s[0] == s[1] {
			return s
		} else {
			return s[0:1]
		}
	}
	dp := make([][]bool, len(s))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(s))
	}

	res_length := 1
	res_string := s[0:1]

	for i := 0; i < len(s); i++ {
		dp[i][i] = true
		if i+1 < len(s) {
			dp[i][i+1] = s[i] == s[i+1]
			if dp[i][i+1] && res_length < 2 {
				res_length = 2
				res_string = s[i : i+2]
			}
		}
	}

	for i := 0; i < len(s); i++ {
		x, y := i-1, i+1
		for x >= 0 && (x+1) < len(s) && (y-1) >= 0 && y < len(s) {
			dp[x][y] = s[x] == s[y] && dp[x+1][y-1]
			if dp[x][y] && y-x+1 > res_length {
				res_length = y - x + 1
				res_string = s[x : y+1]
			}
			x--
			y++
		}

		x, y = i-1, i+1+1
		for x >= 0 && (x+1) < len(s) && (y-1) >= 0 && y < len(s) {
			dp[x][y] = s[x] == s[y] && dp[x+1][y-1]
			if dp[x][y] && y-x+1 > res_length {
				res_length = y - x + 1
				res_string = s[x : y+1]
			}
			x--
			y++
		}
	}

	return res_string
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maketree() *TreeNode {
	n1 := &TreeNode{Val: 1}
	n5 := &TreeNode{Val: 5}
	n3 := &TreeNode{Val: 3}
	n4 := &TreeNode{Val: 4}
	n10 := &TreeNode{Val: 10}
	n6 := &TreeNode{Val: 6}
	n9 := &TreeNode{Val: 9}
	n2 := &TreeNode{Val: 2}
	n1.Left = n5
	n1.Right = n3
	n5.Right = n4
	n3.Right = n10
	n3.Left = n6
	n4.Left = n9
	n4.Right = n2
	return n1
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func makelist() *ListNode {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n5 := &ListNode{Val: 5}
	n6 := &ListNode{Val: 6}
	n7 := &ListNode{Val: 7}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	n5.Next = n6
	n6.Next = n7
	return n1
}

func showlist(l *ListNode) {
	for l != nil {
		fmt.Printf("%d ", l.Val)
		l = l.Next
	}
	fmt.Println("")
}
