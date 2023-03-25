/*
 * @lc app=leetcode.cn id=110 lang=golang
 *
 * [110] 平衡二叉树
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

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	// 求出左右子树的深度，如果不平衡，则返回
	if abs(dfs(root.Left), dfs(root.Right)) > 1 {
		return false
	}
	// 如果左右子树平衡，则递归左右子树
	return isBalanced(root.Left) && isBalanced(root.Right)
}
func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(dfs(root.Left), dfs(root.Right)) + 1
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

// @lc code=end
