/*
 * @lc app=leetcode.cn id=111 lang=golang
 *
 * [111] 二叉树的最小深度
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	depth := 0
	if root == nil {
		return depth
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	for len(queue) > 0 {
		depth++
		n := len(queue)
		for i := 0; i < n; i++ {
			cur := queue[0]
			queue = queue[1:]
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
			if cur.Left == nil && cur.Right == nil {
				return depth
			}
		}
	}
	return depth
}
// @lc code=end

