/*
 * @lc app=leetcode.cn id=116 lang=golang
 *
 * [116] 填充每个节点的下一个右侧节点指针
 */
package main
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}
// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

// 核心是需要先加右节点，再加左节点，同时，需要知道前一个节点位置
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	queue := []*Node{}
	queue = append(queue, root)
	for len(queue) > 0 {
		n := len(queue)
		var preNode *Node
		for i := 0; i < n; i++ {
			cur := queue[0]
			queue = queue[1:]
			if cur.Left != nil {
				queue = append(queue, cur.Right)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Left)
			}
			if i == 0 {
				cur.Next = nil
			} else {
				cur.Next = preNode
			}
			preNode = cur
		}
	}
	return root
}

// @lc code=end
