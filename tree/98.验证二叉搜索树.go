/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
 */
package main

import (
	"fmt"
	"math"
)

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
//  ====================
// 解法一：纯递归
// 根节点应该大于左子树的所有节点，小于右子树的所有节点
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	lm, rm := findMax(root.Left), findMin(root.Right)

	if lm >= root.Val || root.Val >= rm {
		return false
	}
	return isValidBST(root.Left) && isValidBST(root.Right)
}
func findMin(root *TreeNode) int {
	if root == nil {
		return math.MaxInt
	}
	return min(root.Val, findMin(root.Left), findMin(root.Right))
}
func findMax(root *TreeNode) int {
	if root == nil {
		return math.MinInt
	}
	return max(root.Val, findMax(root.Left), findMax(root.Right))
}
func max(a, b, c int) int {
	MaxV := a
	if b > MaxV {
		MaxV = b
	}
	if c > MaxV {
		MaxV = c
	}
	return MaxV
}
func min(a, b, c int) int {
	minV := a
	if b < minV {
		minV = b
	}
	if c < minV {
		minV = c
	}
	return minV
}

// ================================
// 解法二：中序遍历后判断是否递增
// var (
//     order []int
// )
// func isValidBST(root *TreeNode) bool {
//     order = []int{}
//     dfs(root)
//     for i := 1; i < len(order); i++ {
//         if order[i] <= order[i-1] {
//             return false
//         }
//     }
//     return true
// }
// func dfs(root *TreeNode){
//     if root == nil {
//         return
//     }
//     dfs(root.Left)
//     order = append(order, root.Val)
//     dfs(root.Right)
// }
// @lc code=end

func main() {
	root := &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 3},
	}
	fmt.Println(isValidBST(root))
}
