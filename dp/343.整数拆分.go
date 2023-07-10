/*
 * @lc app=leetcode.cn id=343 lang=golang
 *
 * [343] 整数拆分
 */

// @lc code=start
func integerBreak(n int) int {
	// 当 i≥2 时，假设对正整数 i 拆分出的第一个正整数是 j（1≤j<i），则有以下两种方案：
	// 将 i 拆分成 j 和 i−j 的和，且 i−j 不再拆分成多个正整数，此时的乘积是 j×(i−j)；
	// 将 i 拆分成 j 和 i−j 的和，且 i−j 继续拆分成多个正整数，此时的乘积是 j×dp[i−j]。
	// 因此，当 j 固定时，有 dp[i]=max(j×(i−j),j×dp[i−j])
    dp := [60]int{}
    dp[2] = 1
    for i := 3; i <= n; i++ {
        for j := 1; j < i; j++ {
            // 在之前的for循环中，就存在一个最大的dp值，所以需要进行对比
            dp[i] = max(dp[i], max(j * dp[i-j], j *(i-j)))
        }
    }
    return dp[n]
}
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
// @lc code=end

