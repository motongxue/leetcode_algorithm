/*
 * @lc app=leetcode.cn id=637 lang=golang
 *
 * [637] 二叉树的层平均值
 */
package main
// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// ==============
// 主要是需要注意float和int类型间的转换
func averageOfLevels(root *TreeNode) []float64 {
	result := []float64{}
	if root == nil {
		return result
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	for len(queue) > 0 {
		n := len(queue)
		tmpSum := 0
		for i := 0; i < n; i++ {
			tmp := queue[0]
			tmpSum += tmp.Val
			queue = queue[1:]
			if tmp.Left != nil {
				queue = append(queue, tmp.Left)
			}
			if tmp.Right != nil {
				queue = append(queue, tmp.Right)
			}
		}
		result = append(result, float64(tmpSum)/float64(n))
	}
	return result
}
// @lc code=end

