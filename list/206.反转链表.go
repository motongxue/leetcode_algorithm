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
 // 递归写法
 func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    ret := reverseList(head.Next)
    head.Next.Next = head
    head.Next = nil
    return ret
}

// 迭代写法
// O	O  ->  O  ->  O
// ↑	↑	   ↑	  ↑
// prev cur    next
//
// O <- O      O  ->  O
// ↑	↑	   ↑	  ↑
// 		prev   cur    next
func reverseList(head *ListNode) *ListNode {
    var prev *ListNode
    cur := head
    for cur != nil {
        next := cur.Next
        cur.Next = prev
        prev = cur
        cur = next
    }
    return prev
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
