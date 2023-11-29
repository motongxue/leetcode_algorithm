/*
 * @lc app=leetcode.cn id=309 lang=golang
 *
 * [309] 最佳买卖股票时机含冷冻期
 */

// @lc code=start
/* 
dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
dp[i][1] = max(dp[i-1][1], dp[i-2][0] - prices[i])
解释：第 i 天选择 buy 的时候，要从 i-2 的状态转移，而不是 i-1 。
 */
func maxProfit(prices []int) int {
    n := len(prices)
    dp := make([][2]int, n)

    for i := 0; i < n; i++ {
        if i-1 == -1 {
            dp[i][0] = 0
            dp[i][1] = -prices[i]
            continue
        }
        if i-2 == -1 {
            dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
            dp[i][1] = max(dp[i-1][1], - prices[i])
            continue
        }
        dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
        dp[i][1] = max(dp[i-1][1], dp[i-2][0] - prices[i])
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

