/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
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
	mp map[int]int
)

func buildTree(preorder []int, inorder []int) *TreeNode {
	mp = make(map[int]int)
	for i, v := range inorder {
		mp[v] = i
	}
	return build(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1)
}
func build(preorder []int, p_l, p_r int, inorder []int, i_l, i_r int) *TreeNode {
	if p_l > p_r || i_l > i_r || p_l < 0 || i_l < 0 {
		return nil
	}
	r := preorder[p_l]
	r_idx := mp[r]
	// preorder: [p_l+1,r_idx-i_l+p_l] [r_idx+1-i_l+p_l,p_r]
	// inorder: [i_l,r_idx-1] [r_idx+1,i_r]
	return &TreeNode{
		Val:   r,
		Left:  build(preorder, p_l+1, r_idx-i_l+p_l, inorder, i_l, r_idx-1),
		Right: build(preorder, r_idx+1-i_l+p_l, p_r, inorder, r_idx+1, i_r),
	}
}

// @lc code=end

