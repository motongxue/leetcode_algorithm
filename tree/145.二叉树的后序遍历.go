/*
 * @lc app=leetcode.cn id=145 lang=golang
 *
 * [145] 二叉树的后序遍历
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
func postorderTraversal(root *TreeNode) []int {
	result := []int{}
	post(root, &result)
	return result
}
func post(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	post(root.Left, result)
	post(root.Right, result)
	(*result) = append((*result), root.Val)
}

// @lc code=end

