/*
 * @lc app=leetcode.cn id=538 lang=golang
 *
 * [538] 把二叉搜索树转换为累加树
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
// 此时不能通过sum递归传值的方式进行，而是需要采用全局变量
var sum int

func convertBST(root *TreeNode) *TreeNode {
	sum = 0
	dfs(root)
	return root
}
func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Right)
	sum += root.Val
	root.Val = sum
	dfs(root.Left)
}

// @lc code=end

