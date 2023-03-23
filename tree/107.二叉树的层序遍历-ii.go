/*
 * @lc app=leetcode.cn id=107 lang=golang
 *
 * [107] 二叉树的层序遍历 II
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
func levelOrderBottom(root *TreeNode) [][]int {
	result := [][]int{}
	var f func(root *TreeNode,depth int)
	f = func(root *TreeNode, depth int){
		if root == nil {
			return
		}
		if len(result)==depth {
			result = append(result,[]int{})
		}
		f(root.Left,depth+1)
		f(root.Right,depth+1)
		result[depth] = append(result[depth],root.Val)
	}
	f(root,0)
	i,j := 0,len(result)-1
	for i<j {
		result[i],result[j] = result[j],result[i]
		i++
		j--
	}
	return result
}
// @lc code=end

