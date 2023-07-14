/*
 * @lc app=leetcode.cn id=416 lang=golang
 *
 * [416] 分割等和子集
 */

// @lc code=start
// ========================================================
// 使用传统的01背包定义
// func canPartition(nums []int) bool {
// 	n := len(nums)
// 	value := make([]int, n)
// 	weight := make([]int, n)
// 	copy(value, nums)
// 	copy(weight, nums)
// 	sum := 0
// 	for _, v := range nums {
// 		sum += v
// 	}
// 	if sum%2 == 1 {
// 		return false
// 	}
// 	bagWeight := sum/2
// 	dp := make([]int, bagWeight+1)
// 	// 实现01背包
// 	for i := 1; i < n; i++ {
// 		for j := bagWeight; j >= weight[i]; j-- {
// 			dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
// 		}
// 	}
// 	return dp[bagWeight] == bagWeight
// }
// // max函数
// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// ========================================================
// 利用变量值相同简化
func canPartition(nums []int) bool {
    sum := 0
    for _, v := range nums {
        sum += v
    }
    // 如果 nums 的总和为奇数则不可能平分成两个子集
    if sum % 2 == 1 {
        return false
    }
    target := sum/2
    // 为什么dp的长度是target？
    // dp的长度是所能容纳的重量
    // dp的含义：dp[j]为容量为j的背包所背的最大价值
    dp := make([]int, target+1)
    for i := 0; i < len(nums); i++{
        for j := target; j >= nums[i]; j-- {
            dp[j] = max(dp[j], dp[j- nums[i]] + nums[i])
        }
    }
    return dp[target] == target
}
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
// @lc code=end

