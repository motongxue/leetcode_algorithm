/*
 * @lc app=leetcode.cn id=106 lang=golang
 *
 * [106] 从中序与后序遍历序列构造二叉树
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
	mp_in map[int]int
)
func buildTree(inorder []int, postorder []int) *TreeNode {
	mp_in = make(map[int]int)
	for i, v := range inorder {
		mp_in[v] = i
	}
	return build(inorder, 0, len(inorder)-1, postorder, 0, len(postorder)-1)
}
func build(inorder []int, i_l, i_r int, postorder []int, p_l, p_r int) *TreeNode {
	if i_l > i_r || p_l > p_r {
		return nil
	}
	r := postorder[p_r]
	in_r_idx := mp_in[r]
	// 利用map来找到根节点在中序的位置 -> 找到中序的切分点
	// 利用中序数组大小一定是和后序数组的大小相同的，从而通过长度 -> 找到后序数组的切分位置
	// inorder: [i_l,in_r_idx-1] [in_r_idx+1,i_r]
	// postorder: [p_l:p_l+in_r_idx-i_l] [p_l+in_r_idx-i_l:p_r]
	root := &TreeNode{
		Val:   r,
		Left:  build(inorder, i_l, in_r_idx-1, postorder, p_l, p_l+in_r_idx-i_l-1),
		Right: build(inorder, in_r_idx+1, i_r, postorder, p_l+in_r_idx-i_l, p_r-1),
	}
	return root
}

// @lc code=end
