/*
 * @lc app=leetcode.cn id=701 lang=golang
 *
 * [701] 二叉搜索树中的插入操作
 */

// @lc code=start
// =================================================
// 自己的思路：找到插入的节点位置，直接让其父节点指向插入节点，结束递归
// func insertIntoBST(root *TreeNode, val int) *TreeNode {
// 	if root == nil {
// 		return &TreeNode{Val: val}
// 	}
// 	dfs(root, val)
// 	return root
// }
// func dfs(root *TreeNode, val int) {
// 	if root == nil {
// 		return
// 	}
// 	if root.Left == nil && root.Val > val {
// 		root.Left = &TreeNode{Val: val}
// 	} else if root.Right == nil && root.Val < val {
// 		root.Right = &TreeNode{Val: val}
// 	}
// 	if root.Val < val {
// 		dfs(root.Right, val)
// 	} else {
// 		dfs(root.Left, val)
// 	}
// }
// =============================================
// 别人的思路：原地修改
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val: val,
		}
	}
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	}
	if root.Val < val {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}

// @lc code=end

