/*
 * @lc app=leetcode.cn id=429 lang=golang
 *
 * [429] N 叉树的层序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
// 核心是从Children中添加for循环读取child并加入到队列中
func levelOrder(root *Node) [][]int {
    result := [][]int{}
	if root == nil {
		return result
	}
	queue := []*Node{}
	queue = append(queue, root)
	for len(queue) > 0 {
		n := len(queue)
		tmp := []int{}
		for i := 0; i < n; i++ {
			cur := queue[0]
			queue = queue[1:]
			tmp = append(tmp, cur.Val)
			for _,child := range cur.Children {
				queue = append(queue, child)
			}
		}
		result = append(result, tmp)
	}
	return result
}
// @lc code=end

