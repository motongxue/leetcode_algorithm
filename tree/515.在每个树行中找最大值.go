/*
 * @lc app=leetcode.cn id=515 lang=golang
 *
 * [515] 在每个树行中找最大值
 */
package main

import "math"

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	for len(queue) > 0 {
		n := len(queue)
		tmpMax := math.MinInt
		for i := 0; i < n; i++ {
			tmp := queue[0]
			tmpMax = max(tmpMax,tmp.Val)
			queue = queue[1:]
			if tmp.Left != nil {
				queue = append(queue, tmp.Left)
			}
			if tmp.Right != nil {
				queue = append(queue, tmp.Right)
			}
		}
		result = append(result, tmpMax)
	}
	return result
}
func max(a,b int) int{
	if a > b {
		return a
	}
	return b
}
// @lc code=end

