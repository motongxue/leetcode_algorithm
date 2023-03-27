/*
 * @lc app=leetcode.cn id=501 lang=golang
 *
 * [501] 二叉搜索树中的众数
 */

// @lc code=start
// 这种主要是针对普通二叉树，对于二叉搜索树，没有利用其特性
/* var (
	res map[int]int
)

func findMode(root *TreeNode) []int {
	res = make(map[int]int)
	dfs(root)
	result := []int{}
	resFreq := 0
	for k, v := range res {
		if v > resFreq {
			result = []int{}
			result = append(result, k)
			resFreq = v
		} else if v == resFreq {
			result = append(result, k)
		}
	}
	return result
}
func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Left)
	res[root.Val]++
	dfs(root.Right)
} */
// 这种方式利用了二叉搜索树的特性，主要是在递归过程中进行计数
var (
	maxFreq int
	cnt     int
	maxRes  []int
	pre     *TreeNode
)

func findMode(root *TreeNode) []int {
	maxFreq, cnt = 0, 0
	maxRes = []int{}
	pre = nil
	dfs(root)
	return maxRes
}
func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	dfs(root.Left)
	// 注意是值不相等才cnt=1，不同节点地址肯定不同
	if pre == nil || pre.Val != root.Val {
		cnt = 1
	} else {
		cnt++
	}
	// fmt.Printf("cnt count: %v\n",cnt)
	if cnt > maxFreq {
		maxFreq = cnt
		maxRes = []int{}
		maxRes = append(maxRes, root.Val)
	} else if cnt == maxFreq {
		maxRes = append(maxRes, root.Val)
	}
	pre = root
	dfs(root.Right)
}

// @lc code=end

