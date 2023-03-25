/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
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
//  ============================
//  迭代法
// func maxDepth(root *TreeNode) int {
// 	depth := 0
// 	if root == nil {
// 		return depth
// 	}
// 	queue := []*TreeNode{}
// 	queue = append(queue, root)
// 	for len(queue) > 0 {
// 		n := len(queue)
// 		for i := 0; i < n; i++ {
// 			cur := queue[0]
// 			queue = queue[1:]
// 			if cur.Left != nil {
// 				queue = append(queue, cur.Left)
// 			}
// 			if cur.Right != nil {
// 				queue = append(queue, cur.Right)
// 			}
// 		}
// 		depth++
// 	}
// 	return depth
// }

// ===========================
// 递归法
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

