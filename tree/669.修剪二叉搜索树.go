/*
 * @lc app=leetcode.cn id=669 lang=golang
 *
 * [669] 修剪二叉搜索树
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
var (
	nums []int
)

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return root
	}
	nums = []int{}
	dfs(root)
	for _, v := range nums {
		if low <= v && v <= high {
			continue
		}
		root = trim(root, v)
	}
	return root
}
func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Left)
	nums = append(nums, root.Val)
	dfs(root.Right)
}
func trim(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val == key {
		if root.Left != nil && root.Right == nil {
			return root.Left
		} else if root.Left == nil && root.Right != nil {
			return root.Right
		} else if root.Left == nil && root.Right == nil {
			return nil
		} else {
			cur := root.Right
			for cur.Left != nil {
				cur = cur.Left
			}
			cur.Left = root.Left
			return root.Right
		}
	}
	if root.Val > key {
		root.Left = trim(root.Left, key)
	} else {
		root.Right = trim(root.Right, key)
	}
	return root
}

// @lc code=end

