/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
// 同样可以使用变量代替数组
func climbStairs(n int) int {
    dp := [46]int{1,2}
    for i := 2; i < n; i++ {
        // 当前台阶，要么从前一个台阶跨一步，或者是前两个台阶跨两步来的
        dp[i] = dp[i-1] + dp[i-2]
    }
    return dp[n-1]
}
// @lc code=end

