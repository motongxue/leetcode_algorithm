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
// ======================================
// 解法一：十分笨重，先利用中序遍历获取每个值，然后采用[450] 删除二叉搜索树中的节点中的方法，将节点逐个判断并删除
// var (
// 	nums []int
// )

// func trimBST(root *TreeNode, low int, high int) *TreeNode {
// 	if root == nil {
// 		return root
// 	}
// 	nums = []int{}
// 	dfs(root)
// 	for _, v := range nums {
// 		if low <= v && v <= high {
// 			continue
// 		}
// 		root = trim(root, v)
// 	}
// 	return root
// }
// func dfs(root *TreeNode) {
// 	if root == nil {
// 		return
// 	}
// 	dfs(root.Left)
// 	nums = append(nums, root.Val)
// 	dfs(root.Right)
// }
// func trim(root *TreeNode, key int) *TreeNode {
// 	if root == nil {
// 		return root
// 	}
// 	if root.Val == key {
// 		if root.Left != nil && root.Right == nil {
// 			return root.Left
// 		} else if root.Left == nil && root.Right != nil {
// 			return root.Right
// 		} else if root.Left == nil && root.Right == nil {
// 			return nil
// 		} else {
// 			cur := root.Right
// 			for cur.Left != nil {
// 				cur = cur.Left
// 			}
// 			cur.Left = root.Left
// 			return root.Right
// 		}
// 	}
// 	if root.Val > key {
// 		root.Left = trim(root.Left, key)
// 	} else {
// 		root.Right = trim(root.Right, key)
// 	}
// 	return root
// }
// ==========================================
// 解法二：直接利用范围进行判断
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return root
	}
	// 当节点值比low还小时，则看节点的右子树中是否有不比他小的
	if root.Val < low {
		return trimBST(root.Right, low, high)
	}
	if root.Val > high {
		return trimBST(root.Left, low, high)
	}
	// low <= root <= high 时，递归遍历
	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	return root
}

// @lc code=end

