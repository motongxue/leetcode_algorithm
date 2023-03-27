/*
 * @lc app=leetcode.cn id=235 lang=golang
 *
 * [235] 二叉搜索树的最近公共祖先
 */

// @lc code=start
// =======================================
// 这种方法只是通用的二叉树方法，没有利用二叉搜索树的特性
/* func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	l := lowestCommonAncestor(root.Left,p,q)
	r := lowestCommonAncestor(root.Right,p,q)
	if l != nil && r != nil {
		return root
	}else if l != nil && r == nil {
		return l
	}else if l == nil && r != nil {
		return r
	}else {
		return nil
	}
} */
// ===================================
// 利用二叉搜索树特性求解
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || (p.Val <= root.Val && root.Val <= q.Val) || (q.Val <= root.Val && root.Val <= p.Val) {
		fmt.Println(root)
		return root
	}
	if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	} else {
		return lowestCommonAncestor(root.Left, p, q)
	}
}

// @lc code=end
