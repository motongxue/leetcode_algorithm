/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
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
 func preorderTraversal(root *TreeNode) []int {
	result := []int{}
    pre(root,&result)
    return result
}
func pre(root *TreeNode, result *[]int) {
    if root == nil {
        return
    }
    (*result) = append((*result), root.Val)
    pre(root.Left, result)
    pre(root.Right, result)
}
// @lc code=end

