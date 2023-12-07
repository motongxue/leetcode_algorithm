/*
 * @lc app=leetcode.cn id=82 lang=golang
 *
 * [82] 删除排序链表中的重复元素 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
	// 链表中，构建一个空的头结点是非常有用的
    dummy := &ListNode{0, head}

    cur := dummy
    for cur.Next != nil && cur.Next.Next != nil {
        if cur.Next.Val == cur.Next.Next.Val {
            x := cur.Next.Val
			// 循环判断
            for cur.Next != nil && cur.Next.Val == x {
                cur.Next = cur.Next.Next
            }
        } else {
			// 由于相等时，已经进行了next操作，所以需要单独写一个else来next
            cur = cur.Next
        }
    }
	// 返回实际的链表头
    return dummy.Next
}
// @lc code=end

