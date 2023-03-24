/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
 */
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 // 主要是判断a,b值的时候，需要先判断a，b是否为空，否则出现空指针异常
func isSymmetric(root *TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		return true
	}
	if ((root.Left == nil) && (root.Right != nil)) || ((root.Left != nil) && (root.Right == nil)) || (root.Left.Val != root.Right.Val) {
		return false
	}
	return isSymmetricNodes(root.Left, root.Right)
}
func isSymmetricNodes(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if (a == nil && b != nil) || ((a != nil) && (b == nil)) || (a.Left == nil && b.Right != nil) || (a.Left != nil && b.Right == nil) || (a.Val != b.Val) {
		return false
	}
	return isSymmetricNodes(a.Left, b.Right) && isSymmetricNodes(a.Right, b.Left)
}

// @lc code=end

func main() {
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 4,
			},
		},
	}
	fmt.Println(isSymmetric(root))
}
