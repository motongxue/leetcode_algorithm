/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
/* 
dp[i][1][0] = max(dp[i-1][1][0], dp[i-1][1][1] + prices[i])
dp[i][1][1] = max(dp[i-1][1][1], dp[i-1][0][0] - prices[i]) 
            = max(dp[i-1][1][1], -prices[i])
解释：k = 0 的 base case，所以 dp[i-1][0][0] = 0。

现在发现 k 都是 1，不会改变，即 k 对状态转移已经没有影响了。
可以进行进一步化简去掉所有 k：
dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
dp[i][1] = max(dp[i-1][1], -prices[i])
 */
func maxProfit(prices []int) int {
    n := len(prices)
    dp := make([][2]int, n)
	for i := 0; i < n; i++ {
		if i-1 == -1 {
			dp[0][0] = 0
			dp[0][1] = -prices[0]
			continue
		}
    
        dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
        dp[i][1] = max(dp[i-1][1], -prices[i])
    }
    return dp[n-1][0]
}
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
// @lc code=end

