/*
 * @lc app=leetcode.cn id=494 lang=golang
 *
 * [494] 目标和
 */

// @lc code=start
func findTargetSumWays(nums []int, target int) int {
    // left - right = target
    // left + right = sum
    // left = (target + sum)/2
    sum, n := 0, len(nums)
    for _, v := range nums {
        sum += v
    }
    if (target + sum) % 2 == 1 || abs(target) > sum {
        return 0
    }
    bagSize := (target + sum)/2
    // dp[j] 表示：填满j（包括j）这么大容积的包，有dp[j]种方法
    dp := make([]int, bagSize+1)
    // 表示从前0个数字中选出若干个数字使得其和为0的方案数为1，即「空集合」不选任何数字即可得到0。
    dp[0] = 1
    for i := 0; i < n; i++ {
        for j := bagSize; j >= nums[i]; j-- {
            dp[j] += dp[j-nums[i]]
        }
    }
    return dp[bagSize]
}
func abs(a int) int {
    if a >= 0 {
        return a
    }
    return -a
}
// @lc code=end

