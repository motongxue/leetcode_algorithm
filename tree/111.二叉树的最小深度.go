/*
 * @lc app=leetcode.cn id=111 lang=golang
 *
 * [111] 二叉树的最小深度
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
// ==================
// 迭代法
// func minDepth(root *TreeNode) int {
// 	depth := 0
// 	if root == nil {
// 		return depth
// 	}
// 	queue := []*TreeNode{}
// 	queue = append(queue, root)
// 	for len(queue) > 0 {
// 		depth++
// 		n := len(queue)
// 		for i := 0; i < n; i++ {
// 			cur := queue[0]
// 			queue = queue[1:]
// 			if cur.Left != nil {
// 				queue = append(queue, cur.Left)
// 			}
// 			if cur.Right != nil {
// 				queue = append(queue, cur.Right)
// 			}
// 			if cur.Left == nil && cur.Right == nil {
// 				return depth
// 			}
// 		}
// 	}
// 	return depth
// }
// =====================
// 递归法
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// TODO 只有当左右子树均为空时，才是最小深度
	if root.Left == nil && root.Right != nil {
		return minDepth(root.Right) + 1
	} else if root.Left != nil && root.Right == nil {
		return minDepth(root.Left) + 1
	} else {
		// 出现两种情况：
		//     1. left、right都为nil -> 由于上面的root=nil时，返回0，故正确
		//     2. left、right都不为nil -> 则取最小值
		return min(minDepth(root.Left), minDepth(root.Right)) + 1
	}
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
