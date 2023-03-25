/*
 * @lc app=leetcode.cn id=257 lang=golang
 *
 * [257] 二叉树的所有路径
 */
package main

import (
	"strconv"
)
type TreeNode struct {
    Val int
    Left *TreeNode
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
// TODO　难点在于字符串的处理以及匿名内部函数的定义
// path = path + strconv.Itoa(root.Val) + "->"
func binaryTreePaths(root *TreeNode) []string {
	result := []string{}
	var path string
	var f func(root *TreeNode, path string)
	f = func(root *TreeNode, path string) {
		// 重点在于递归条件的终止
		if root.Left == nil && root.Right == nil {
			v := path + strconv.Itoa(root.Val)
			result = append(result, v)
			return
		}
		path = path + strconv.Itoa(root.Val) + "->"
		// 为什么要先判断?由终止条件所限定了
		if root.Left != nil {
			f(root.Left, path)
		}
		if root.Right != nil {
			f(root.Right, path)
		}
	}
	f(root, path)
	return result
}

// @lc code=end
