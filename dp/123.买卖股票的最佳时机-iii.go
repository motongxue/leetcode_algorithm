/*
 * @lc app=leetcode.cn id=123 lang=golang
 *
 * [123] 买卖股票的最佳时机 III
 */

// @lc code=start
/* 
原始的状态转移方程，没有可化简的地方
dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])
 */
func maxProfit(prices []int) int {
    n := len(prices)
    dp := make([][3][2]int, n)
    for i := 0; i < n; i++ {
        for k := 1; k <= 2; k++ {
            if i-1 == -1 {
                dp[i][k][0] = 0
                dp[i][k][1] = -prices[i]
                continue
            }
            dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
            dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])
        }
    }
    return dp[n-1][2][0]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
// @lc code=end

