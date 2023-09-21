/*
 * @lc app=leetcode.cn id=257 lang=golang
 *
 * [257] 二叉树的所有路径
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
//  =====================================================
// 0：该节点无覆盖
// 1：本节点有摄像头
// 2：本节点有覆盖

// case 1: 左右节点都有覆盖
// case 2: 左右节点存在一个无覆盖
// case 3: 左右节点存在一个为摄像头
// case 4: 判断root节点，root节点是否也需要被监视
// TODO 重做
func minCameraCover(root *TreeNode) int {
	res := 0
	var findMin func(root *TreeNode) int
	findMin = func(root *TreeNode) int {
		if root == nil {
			return 2
		}
		left, right := findMin(root.Left), findMin(root.Right)
		if left == 2 && right == 2 {
			// case 1
			return 0
		} else if left == 0 || right == 0 {
			// case 2
			res++
			return 1
		} else if left == 1 || right == 1 {
			// case 3
			return 2
		}
		return -1
	}
	// case 4
	if findMin(root) == 0 {
		res++
	}
	return res
}

// @lc code=end

