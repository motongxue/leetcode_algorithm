/*
 * @lc app=leetcode.cn id=509 lang=golang
 *
 * [509] 斐波那契数
 */

// @lc code=start
// func fib(n int) int {
// 	dp := [31]int{}
// 	dp[0] = 0
// 	dp[1] = 1
// 	for i := 2; i <= n; i++ {
// 		dp[i] = dp[i-1] + dp[i-2]
// 	}
// 	return dp[n]
// }
func fib(n int) int {
	if n < 2 {
		return n
	}
	// 定义变量代替数组节约空间
	fn := 0
	fn_2, fn_1 := 0, 1
	for i := 2; i <= n; i++ {
		fn = fn_1 + fn_2
		fn_2 = fn_1
		fn_1 = fn
	}
	return fn
}
// @lc code=end

