/*
 * @lc app=leetcode.cn id=654 lang=golang
 *
 * [654] 最大二叉树
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
func constructMaximumBinaryTree(nums []int) *TreeNode {
	return build(0, len(nums)-1, nums)
}
func build(l, r int, nums []int) *TreeNode {
	if l > r {
		return nil
	}
	MaxV := math.MinInt
	MaxIdx := -1
	for i := l; i <= r; i++ {
		if nums[i] > MaxV {
			MaxV = nums[i]
			MaxIdx = i
		}
	}
	return &TreeNode{
		Val:   MaxV,
		Left:  build(l, MaxIdx-1, nums),
		Right: build(MaxIdx+1, r, nums),
	}
}

// @lc code=end

