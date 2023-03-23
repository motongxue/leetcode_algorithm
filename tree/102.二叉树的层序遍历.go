/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
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
//  ================
// 采用递归 + deep变量
// func levelOrder(root *TreeNode) [][]int {
// 	result := [][]int{}
// 	deepFunc(root, 0, &result)
// 	return result
// }
// func deepFunc(root *TreeNode, deep int, result *[][]int) {
// 	if root == nil {
// 		return
// 	}
// 	if len((*result)) == deep {
// 		(*result) = append((*result), []int{})
// 	}
// 	(*result)[deep] = append((*result)[deep], root.Val)
// 	deepFunc(root.Left, deep+1, result)
// 	deepFunc(root.Right, deep+1, result)
// }

func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	n := len(queue)
	tmp := []int{}
	for len(queue) > 0 {
		tmp = append(tmp, queue[0].Val)
		if queue[0].Left != nil {
			queue = append(queue, queue[0].Left)
		}
		if queue[0].Right != nil {
			queue = append(queue, queue[0].Right)
		}
		queue = queue[1:]
		n--
		if n == 0 {
			result = append(result, tmp)
			tmp = []int{}
			n = len(queue)
		}
	}
	return result
}

// @lc code=end
func main() {
	root := &TreeNode{
		Val:   3,
		Left:  &TreeNode{Val: 9},
		Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
	}

	fmt.Println(levelOrder(root))
}
