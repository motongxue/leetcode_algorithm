/*
 * @lc app=leetcode.cn id=746 lang=golang
 *
 * [746] 使用最小花费爬楼梯
 */

// @lc code=start
func minCostClimbingStairs(cost []int) int {
    // 到达楼梯顶部时，要么从n-1或n-2中上来
    dp := [1001]int{}
    dp[0], dp[1] = cost[0], cost[1]
    n := len(cost)
    for i := 2; i < n; i++ {
        // 即代价要么从i-1中来，要么从i-2中来
        dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
    }
    // 若循环时使用：i<=n，则cost[n]=0，可以直接return dp[n]
    return min(dp[n-1], dp[n-2])
}
func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}
// @lc code=end

