/*
 * @lc app=leetcode.cn id=404 lang=golang
 *
 * [404] 左叶子之和
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
//  ================
// 通过dir变量记录是左递归还是右递归
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	var f func(root *TreeNode, dir int)
	f = func(root *TreeNode, dir int) {
		// 此时为叶子节点时候才记录
		if root.Left == nil && root.Right == nil {
			if dir == 0 {
				res += root.Val
			}
			return
		}
		if root.Left != nil {
			f(root.Left, 0)
		}
		if root.Right != nil {
			f(root.Right, 1)
		}
	}
	f(root, 1)
	return res
}

// =======================
// 别人的思路：通过root.Left.Left 来判断
// func sumOfLeftLeaves(root *TreeNode) int {
//     if root == nil {
//         return 0
//     }
//     leftValue := sumOfLeftLeaves(root.Left)   // 左
//     if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
//         leftValue = root.Left.Val             // 中
//     }
//     rightValue := sumOfLeftLeaves(root.Right) // 右
//     return leftValue + rightValue
// }
// @lc code=end

