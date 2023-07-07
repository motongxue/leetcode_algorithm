/*
 * @lc app=leetcode.cn id=63 lang=golang
 *
 * [63] 不同路径 II
 */

// @lc code=start
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    dp := [101][101]int{}
    m, n := len(obstacleGrid), len(obstacleGrid[0])
    // 若边界上为障碍物，则不能通行
    for i := 0; i < m && obstacleGrid[i][0] == 0; i++ {
        dp[i][0] = 1
    }
    for j := 0; j < n && obstacleGrid[0][j] == 0; j++ {
        dp[0][j] = 1
    }
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            // 若道路上为障碍物，则不能通行
            if obstacleGrid[i][j] == 1 {
                continue
            }
            dp[i][j] = dp[i-1][j] + dp[i][j-1]
        }
    }
    return dp[m-1][n-1]
}
// @lc code=end

