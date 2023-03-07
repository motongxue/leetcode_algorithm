/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	dummyHead := &ListNode{
		Next: nil,
	}
	reverse(head, dummyHead)
	return dummyHead.Next
}

func reverse(cur, result *ListNode) {
	if cur != nil {
		reverse(cur.Next, result)
		tmp := result
		for tmp.Next != nil {
			tmp = tmp.Next
		}
		tmp.Next = &ListNode{
			Val: cur.Val,
			Next: nil,
		}
	}
}

// @lc code=end

func main() {

	h5 := &ListNode{
		Val:  5,
		Next: nil,
	}
	h4 := &ListNode{
		Val:  4,
		Next: h5,
	}
	h3 := &ListNode{
		Val:  3,
		Next: h4,
	}
	h2 := &ListNode{
		Val:  2,
		Next: h3,
	}
	h1 := &ListNode{
		Val:  1,
		Next: h2,
	}
	fmt.Println(reverseList(h1))
}
