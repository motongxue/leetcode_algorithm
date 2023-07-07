/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 */

// @lc code=start
func uniquePaths(m int, n int) int {
    dp := [101][101]int{}
    // 对边界赋1
    for i := 0; i < m; i++ {
        dp[i][0] = 1
    }
    for j := 0; j < n; j++ {
        dp[0][j] = 1
    }
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            // 只能向下、向右走
            dp[i][j] = dp[i-1][j] + dp[i][j-1]
        }
    }
    return dp[m-1][n-1]
}
// @lc code=end

