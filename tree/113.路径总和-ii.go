/*
 * @lc app=leetcode.cn id=113 lang=golang
 *
 * [113] 路径总和 II
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
func pathSum(root *TreeNode, targetSum int) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}
	var dfs func(root *TreeNode, targetSum int, path []int)
	dfs = func(root *TreeNode, targetSum int, path []int) {
		targetSum -= root.Val
		path = append(path, root.Val)

		if root.Left == nil && root.Right == nil {
			if targetSum == 0 {
				// TODO 若直接采用result = append(result, path) 则后面对path的修改会影响result中的结果
				tmp := []int{}
				tmp = append(tmp, path...)
				result = append(result, tmp)
			}
			return
		}
		if root.Left != nil {
			dfs(root.Left, targetSum, path)
		}
		if root.Right != nil {
			dfs(root.Right, targetSum, path)
		}
	}
	dfs(root, targetSum, []int{})
	return result
}

// @lc code=end

func main() {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 13,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 5,
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
	}
	fmt.Println(pathSum(root, 22))
}
