/*
 * @lc app=leetcode.cn id=160 lang=golang
 *
 * [160] 相交链表
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
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA := 0
	lenB := 0
	curA := headA
	curB := headB
	for curA != nil {
		lenA++
		curA = curA.Next
	}
	for curB != nil {
		lenB++
		curB = curB.Next
	}

	curA = headA
	curB = headB
	gap := lenA - lenB
	if gap < 0 {
		tmp := curA
		curA = curB
		curB = tmp
		gap = -gap
	}
	for i := 0; i < gap; i++ {
		curA = curA.Next
	}
	for curA != nil {

		if curA == curB {
			return curA
		}
		curA = curA.Next
		curB = curB.Next
	}
	return nil
}

// @lc code=end

func main() {
	h6 := &ListNode{
		Val:  5,
		Next: nil,
	}
	h5 := &ListNode{
		Val:  4,
		Next: h6,
	}
	h4 := &ListNode{
		Val:  8,
		Next: h5,
	}
	h3 := &ListNode{
		Val:  1,
		Next: h4,
	}
	h2 := &ListNode{
		Val:  4,
		Next: h3,
	}
	h66 := &ListNode{
		Val:  5,
		Next: nil,
	}
	h55 := &ListNode{
		Val:  4,
		Next: h66,
	}
	h44 := &ListNode{
		Val:  8,
		Next: h55,
	}
	h33 := &ListNode{
		Val:  1,
		Next: h44,
	}
	h22 := &ListNode{
		Val:  6,
		Next: h33,
	}
	h11 := &ListNode{
		Val:  5,
		Next: h22,
	}
	fmt.Println(getIntersectionNode(h11, h2))
}
