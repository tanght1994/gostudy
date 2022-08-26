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
	fmt.Println("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	a := []int{5, 1, 6, 2, 8, 3, 4}
	fmt.Println(a)
}

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// func maketree() *TreeNode {
// 	n1 := &TreeNode{Val: 1}
// 	n5 := &TreeNode{Val: 5}
// 	n3 := &TreeNode{Val: 3}
// 	n4 := &TreeNode{Val: 4}
// 	n10 := &TreeNode{Val: 10}
// 	n6 := &TreeNode{Val: 6}
// 	n9 := &TreeNode{Val: 9}
// 	n2 := &TreeNode{Val: 2}
// 	n1.Left = n5
// 	n1.Right = n3
// 	n5.Right = n4
// 	n3.Right = n10
// 	n3.Left = n6
// 	n4.Left = n9
// 	n4.Right = n2
// 	return n1
// }

// func showtree(root *TreeNode) {
// 	if root == nil {
// 		return
// 	}
// 	queue := []*TreeNode{root}
// 	one_level := []*TreeNode{}
// 	all_level := [][]*TreeNode{}
// 	level := 0
// 	for {
// 		if len(one_level) == int(math.Pow(2, float64(level))) {
// 			all_level = append(all_level, one_level)
// 			one_level = one_level[:0]
// 			level++
// 		}
// 		cur := queue[0]
// 		if cur == nil {
// 			queue = append(queue, nil, nil)
// 			one_level = append(one_level, nil, nil)
// 		} else {
// 			queue = append(queue, cur.Left, cur.Right)
// 			one_level = append(one_level, cur.Left, cur.Right)
// 		}
// 		queue = queue[1:]
// 	}
// }

// func merge(head1, head2 *ListNode) *ListNode {

// }

// func sort(head, tail *ListNode) *ListNode {

// }

// func sortList(head *ListNode) *ListNode {
// 	return sort(head, nil)
// }

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// func makelist() *ListNode {
// 	n1 := &ListNode{Val: 1}
// 	n2 := &ListNode{Val: 2}
// 	n3 := &ListNode{Val: 3}
// 	n4 := &ListNode{Val: 4}
// 	n5 := &ListNode{Val: 5}
// 	n6 := &ListNode{Val: 6}
// 	n7 := &ListNode{Val: 7}
// 	n1.Next = n2
// 	n2.Next = n3
// 	n3.Next = n4
// 	n4.Next = n5
// 	n5.Next = n6
// 	n6.Next = n7
// 	return n1
// }

// func showlist(l *ListNode) {
// 	for l != nil {
// 		fmt.Printf("%d ", l.Val)
// 		l = l.Next
// 	}
// 	fmt.Println("")
// }
