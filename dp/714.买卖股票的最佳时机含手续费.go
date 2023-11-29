/*
 * @lc app=leetcode.cn id=714 lang=golang
 *
 * [714] 买卖股票的最佳时机含手续费
 */

// @lc code=start
// 在卖出去时顺便将手续费减掉即可
func maxProfit(prices []int, fee int) int {
    n := len(prices)
    dp := make([][2]int, n)
    for i := 0; i < n; i++ {
        if i-1 == -1 {
            dp[i][0] = 0
            dp[i][1] = -prices[i]
            continue
        }
        dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i] - fee)
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

