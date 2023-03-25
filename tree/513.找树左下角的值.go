/*
 * @lc app=leetcode.cn id=513 lang=golang
 *
 * [513] 找树左下角的值
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
//  ===================================
// 迭代法
// 先加left节点，i=0时为最左元素
//  再利用层次遍历先找到最下层
// func findBottomLeftValue(root *TreeNode) int {
// 	queue := []*TreeNode{}
// 	queue = append(queue, root)
// 	res := root.Val
// 	for len(queue) > 0 {
// 		n := len(queue)
// 		for i := 0; i < n; i++ {
// 			cur := queue[0]
// 			queue = queue[1:]
// 			if i == 0 {
// 				res = cur.Val
// 			}
// 			if cur.Left != nil {
// 				queue = append(queue, cur.Left)
// 			}
// 			if cur.Right != nil {
// 				queue = append(queue, cur.Right)
// 			}
// 		}
// 	}
// 	return res
// }

// =======================
// 递归法
func findBottomLeftValue(root *TreeNode) int {
	resDep, resLeft := 0, root.Val
	var dfs func(root *TreeNode, depth int)
	dfs = func(root *TreeNode, depth int) {
		if root.Left == nil && root.Right == nil {
			// 当第一次出现时记录下来
			if depth > resDep {
				resDep = depth
				resLeft = root.Val
			}
			return
		}
		if root.Left != nil {
			dfs(root.Left, depth+1)
		}
		if root.Right != nil {
			dfs(root.Right, depth+1)
		}
	}
	dfs(root, 0)
	return resLeft
}

// @lc code=end