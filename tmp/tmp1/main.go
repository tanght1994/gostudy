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
	tree := maketree()
	amountOfTime(tree, 3)
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

// func sortList(head *ListNode) *ListNode {

// }

// func sorttt(nums []int) {
// 	for i := 0; i < 100; i++ {
// 		len(nums)
// 	}
// }

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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func amountOfTime(root *TreeNode, start int) int {
	root_to_start := 0
	start_is_root := false
	start_in_left := false
	left_max_deep := 0
	right_max_deep := 0

	if root.Val == start {
		root_to_start = 0
		start_is_root = true
	} else {
		deep := find_target(root.Left, start, 1)
		if deep == -1 {
			start_in_left = false
			deep = find_target(root.Right, start, 1)
		} else {
			start_in_left = true
		}
		root_to_start = deep
	}

	left_max_deep = get_deep(root.Left, 0)
	right_max_deep = get_deep(root.Right, 0)

	if start_is_root {
		return max(left_max_deep, right_max_deep)
	}
	if start_in_left {
		return max(root_to_start+right_max_deep, left_max_deep-root_to_start)
	} else {
		return max(root_to_start+left_max_deep, right_max_deep-root_to_start)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func find_target(root *TreeNode, target, deep int) int {
	if root == nil {
		return -1
	}
	if root.Val == target {
		return deep
	}
	deep1 := find_target(root.Left, target, deep+1)
	deep2 := find_target(root.Right, target, 1)
	return max(deep1, deep2)
}

func get_deep(root *TreeNode, deep int) int {
	if root == nil {
		return deep
	}
	deep1 := get_deep(root.Left, deep+1)
	deep2 := get_deep(root.Right, deep+1)
	return max(deep1, deep2)
}

func merge(head1, head2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	temp, temp1, temp2 := dummyHead, head1, head2
	for temp1 != nil && temp2 != nil {
		if temp1.Val <= temp2.Val {
			temp.Next = temp1
			temp1 = temp1.Next
		} else {
			temp.Next = temp2
			temp2 = temp2.Next
		}
		temp = temp.Next
	}
	if temp1 != nil {
		temp.Next = temp1
	} else if temp2 != nil {
		temp.Next = temp2
	}
	return dummyHead.Next
}

func sort(head, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}

	if head.Next == tail {
		head.Next = nil
		return head
	}

	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}

	mid := slow
	return merge(sort(head, mid), sort(mid, tail))
}

func sortList(head *ListNode) *ListNode {
	return sort(head, nil)
}

func aaaaaa(head1, head2 *ListNode) *ListNode {
	fakehead := &ListNode{}
	tmp1, tmp2, tmp3 := fakehead, head1, head2
	for tmp2 != nil && tmp3 != nil {
		if tmp2.Val < tmp3.Val {
			tmp1.Next = tmp2
			tmp2 = tmp2.Next
		} else {
			tmp1.Next = tmp3
			tmp3 = tmp3.Next
		}
		tmp1 = tmp1.Next
	}
	if tmp2 != nil {
		tmp1.Next = tmp2
	} else if tmp3 != nil {
		tmp1.Next = tmp3
	}
	return fakehead.Next
}
