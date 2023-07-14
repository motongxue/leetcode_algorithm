/*
 * @lc app=leetcode.cn id=96 lang=golang
 *
 * [96] 不同的二叉搜索树
 */

// @lc code=start
func numTrees(n int) int {
	// 初始化dp[0] = 1 防止j=1的时候，导致dp[j-1] = 0
	dp := [20]int{1}

	for i := 1; i <= n; i++ {
		// 从1到n中，将每个节点都挑选来当根节点，左边右边一起得出来后相乘
		// 因此j需要从1到n
		for j := 1; j <= i; j++ {
			// 比如：dp[0], dp[1], dp[2] = 1, 1, 2
			// dp[3] = dp[2] * dp[0] + dp[1] * dp[1] + dp[0] * dp[2]
			// dp[3] = 2*1 + 1*1 + 1*2 = 5
			dp[i] += dp[i-j] * dp[j-1];
		}
	}
	return dp[n]
}
// @lc code=end