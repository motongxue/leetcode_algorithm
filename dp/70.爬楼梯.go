/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
// 同样可以使用变量代替数组
// func climbStairs(n int) int {
//     dp := [46]int{1,2}
//     for i := 2; i < n; i++ {
//         // 当前台阶，要么从前一个台阶跨一步，或者是前两个台阶跨两步来的
//         dp[i] = dp[i-1] + dp[i-2]
//     }
//     return dp[n-1]
// }

// 定义了step，则代表每次可以爬step个楼梯，变成了完全背包问题
func climbStairs(n int) int {
    dp := [46]int{0,1,2}
    step := 2
    for i := 3; i <= n; i++ {
        for j := 1; j <= step; j++ {
            if i - j > 0 {
                dp[i] += dp[i-j]
            }
        }
    }
    return dp[n]
}
// @lc code=end

