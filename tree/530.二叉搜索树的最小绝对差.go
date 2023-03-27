/*
 * @lc app=leetcode.cn id=530 lang=golang
 *
 * [530] 二叉搜索树的最小绝对差
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
/* 
// =================================================
// 由于是二叉搜索树，直接中序遍历即可
var (
	result []int
)
func getMinimumDifference(root *TreeNode) int {
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
} */
// =========================
// 利用pre指针+二叉搜索树特性直接求值
var (
	resMin int
	pre *TreeNode

)
func getMinimumDifference(root *TreeNode) int {
	pre = nil
	resMin = math.MaxInt
	dfs(root)
	return resMin
}
func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Left)
	if pre != nil {
		tmpMin := abs(pre.Val,root.Val)
		if tmpMin < resMin {
			resMin = tmpMin
		}
	}
	pre = root
	dfs(root.Right)
}
func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
} 
// @lc code=end

