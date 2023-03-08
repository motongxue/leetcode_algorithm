/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
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
func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{
		Next: head,
	}
	cur := dummyHead
	for cur.Next != nil && cur.Next.Next != nil {
		tmp1 := cur.Next
		tmp2 := cur.Next.Next
		tmp1.Next = tmp2.Next
		tmp2.Next = tmp1
		cur.Next = tmp2
		cur = cur.Next.Next
	}
	return dummyHead.Next
}
// 递归版本
// func swapPairs(head *ListNode) *ListNode {
//     if head == nil || head.Next == nil {
//         return head
//     }
//     next := head.Next
//     head.Next = swapPairs(next.Next)
//     next.Next = head
//     return next
// }

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
	fmt.Println(swapPairs(h1))
}
