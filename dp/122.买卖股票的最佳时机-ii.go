/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 */

// @lc code=start
/* 
dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])
            = max(dp[i-1][k][1], dp[i-1][k][0] - prices[i])

我们发现数组中的 k 已经不会改变了，也就是说不需要记录 k 这个状态了：
dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
dp[i][1] = max(dp[i-1][1], dp[i-1][0] - prices[i])
 */
func maxProfit(prices []int) int {
    n := len(prices)
    dp := make([][2]int, n)

    for i := 0; i < n; i++ {
		if i == 0 {
			dp[0][0] = 0
			dp[0][1] = -prices[0]
			continue
		}
        dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
        dp[i][1] = max(dp[i-1][1], dp[i-1][0] - prices[i])
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

