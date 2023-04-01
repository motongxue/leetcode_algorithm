/*
 * @lc app=leetcode.cn id=108 lang=golang
 *
 * [108] 将有序数组转换为二叉搜索树
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
// 找到中间点来作为根节点，然后递归左右子数组，重复操作
func sortedArrayToBST(nums []int) *TreeNode {
	l, mid, r := 0, len(nums)/2, len(nums)
	root := &TreeNode{
		Val: nums[mid],
	}
	// 判断时需要先限定条件，否则会导致越界或其他操作
	if mid > l {
		root.Left = sortedArrayToBST(nums[l:mid])
	}
	if r > mid+1 {
		root.Right = sortedArrayToBST(nums[mid+1 : r])
	}
	return root
}

// @lc code=end

