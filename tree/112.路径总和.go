/*
 * @lc app=leetcode.cn id=112 lang=golang
 *
 * [112] 路径总和
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
// 1. 叶子节点
// 2. 利用两个变量bl、br来进行或判断
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	targetSum -= root.Val
	// 如果达到叶子节点，则判断targetSum是否为0
	if root.Left == nil && root.Right == nil {
		return targetSum == 0
	}
	bl := false
	br := false
	// 由于节点的值可以为负
	// 所以不可以利用 targetSum - root.Left.Val >=0 来跳过一些递归过程
	if root.Left != nil {
		bl = hasPathSum(root.Left, targetSum)
	}
	if root.Right != nil {
		br = hasPathSum(root.Right, targetSum)
	}
	return bl || br
}

// @lc code=end
