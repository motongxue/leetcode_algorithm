/*
 * @lc app=leetcode.cn id=783 lang=golang
 *
 * [783] 二叉搜索树节点最小距离
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
var (
	result []int
)

// 由于是二叉搜索树，直接中序遍历即可
func minDiffInBST(root *TreeNode) int {
	res := math.MaxInt
	result = []int{}
	dfs(root)
	for i := 1; i < len(result); i++ {
		tmp := abs(result[i], result[i-1])
		if tmp < res {
			res = tmp
		}
	}
	return res
}
func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Left)
	result = append(result, root.Val)
	dfs(root.Right)
}
func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

// @lc code=end

