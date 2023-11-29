/*
 * @lc app=leetcode.cn id=188 lang=golang
 *
 * [188] 买卖股票的最佳时机 IV
 */

// @lc code=start
// 纯纯的模板题目
func maxProfit(K int, prices []int) int {
    n := len(prices)

    dp := make([][101][2]int, n)
    for i := 0; i < n; i++ {
        for k := 1; k <= K; k++ {
            if i == 0 {
                dp[i][k][0] = 0
                dp[i][k][1] = -prices[i]
                continue
            }
            dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
            dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])
        }
    }
    return dp[n-1][K][0]
}
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
// @lc code=end

