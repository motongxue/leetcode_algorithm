/*
 * @lc app=leetcode.cn id=142 lang=golang
 *
 * [142] 环形链表 II
 */
// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 双指针的解法
// func detectCycle(head *ListNode) *ListNode {
// 	mp := make(map[*ListNode]bool)
// 	for head != nil {
// 		if mp[head] {
// 			return head
// 		}
// 		mp[head] = true
// 		head = head.Next
// 	}
// 	return nil
// }

// 快慢指针解法
func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for true {
		if fast == nil || fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	fast = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}
// @lc code=end

