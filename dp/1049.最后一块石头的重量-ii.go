/*
 * @lc app=leetcode.cn id=1049 lang=golang
 *
 * [1049] 最后一块石头的重量 II
 */

// @lc code=start
func lastStoneWeightII(stones []int) int {
	n := len(stones)
	sum := 0
	for _, v := range stones {
		sum += v
	}
	// 存在精度丢失问题
	bagWeight := sum / 2
	dp := make([]int, bagWeight+1)
	for i := 0; i < n; i++ {
		for j := bagWeight; j >= stones[i]; j-- {
			dp[j] = max(dp[j], dp[j-stones[i]]+stones[i])
		}
	}
	// 实际上就是sum的一半大小的背包容量，能装多少，也就是能被碰撞的数量的一半
	// dp[bagWeight]*2即为能被碰撞数量，那么剩下的即为剩下的石头
	return sum - dp[bagWeight]*2
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

