/*
 * @lc app=leetcode.cn id=222 lang=golang
 *
 * [222] 完全二叉树的节点个数
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
// =============
// 时间复杂度O(n)
// func countNodes(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}
// 	return countNodes(root.Left) + countNodes(root.Right) + 1
// }
// =====================
// 时间复杂度O(log n × log n)
// 利用左子树高度和右子树高度是否相等判断是否为满二叉树
func countNodes(root *TreeNode) int {
    hl, hr := 0, 0
    l, r := root, root
    for l != nil {
        l = l.Left
        hl++
    }
    for r != nil {
        r = r.Right
        hr++
    }
    // 如果左右子树的高度相同，则是一棵满二叉树
    if hl == hr {
        return int(math.Pow(2, float64(hl))) - 1
    }
    // 如果左右高度不同，则按照普通二叉树的逻辑计算
    return 1 + countNodes(root.Left) + countNodes(root.Right)
}
// @lc code=end

