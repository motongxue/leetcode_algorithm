/*
 * @lc app=leetcode.cn id=617 lang=golang
 *
 * [617] 合并二叉树
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
//  =====================
// 自己的版本：直接进行模拟就行，注意当一方节点已经为nil，则不能调用nil.left right指针，否则出现空指针异常
// func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
// 	if root1 == nil && root2 == nil {
// 		return nil
// 	}else if root1 == nil && root2 != nil {
// 		return &TreeNode{
// 			Val: root2.Val,
// 			Left: mergeTrees(nil,root2.Left),
// 			Right: mergeTrees(nil,root2.Right),
// 		}
// 	}else if root1 != nil && root2 == nil {
// 		return &TreeNode{
// 			Val: root1.Val,
// 			Left: mergeTrees(root1.Left,nil),
// 			Right: mergeTrees(nil,root1.Right),
// 		}
// 	}else {
// 		// 都不为nil
// 		return &TreeNode{
// 			Val: root1.Val+root2.Val,
// 			Left: mergeTrees(root1.Left,root2.Left),
// 			Right: mergeTrees(root1.Right,root2.Right),
// 		}
// 	}
// }
// =====================
// 别人的版本：可以省去一些重复操作
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	} else if root2 == nil {
		return root1
	}
	return &TreeNode{
		Val:   root1.Val + root2.Val,
		Left:  mergeTrees(root1.Left, root2.Left),
		Right: mergeTrees(root1.Right, root2.Right),
	}
}

// @lc code=end

