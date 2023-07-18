/*
 * @lc app=leetcode.cn id=494 lang=golang
 *
 * [494] 目标和
 */

// @lc code=start
// ========================================
// 01背包问题
// func findTargetSumWays(nums []int, target int) int {
//     // left - right = target
//     // left + right = sum
//     // left = (target + sum)/2
//     sum, n := 0, len(nums)
//     for _, v := range nums {
//         sum += v
//     }
//     if (target + sum) % 2 == 1 || abs(target) > sum {
//         return 0
//     }
//     bagSize := (target + sum)/2
//     // dp[j] 表示：填满j（包括j）这么大容积的包，有dp[j]种方法
//     dp := make([]int, bagSize+1)
//     // 表示从前0个数字中选出若干个数字使得其和为0的方案数为1，即「空集合」不选任何数字即可得到0。
//     dp[0] = 1
//     for i := 0; i < n; i++ {
//         for j := bagSize; j >= nums[i]; j-- {
//             dp[j] += dp[j-nums[i]]
//         }
//     }
//     return dp[bagSize]
// }
// func abs(a int) int {
//     if a >= 0 {
//         return a
//     }
//     return -a
// }

// ========================================
// 回溯解法
// def backtrack(nums, i):
//     if i == len(nums):
//         if 达到 target:
//             result += 1
//         return

//     for op in { +1, -1 }:
//         选择 op * nums[i]
//         # 穷举 nums[i + 1] 的选择
//         backtrack(nums, i + 1)
//         撤销选择
func findTargetSumWays(nums []int, target int) int {
	res := 0
	var backtrack func(i, sum int)
	backtrack = func(i, sum int) {
		if i == len(nums) {
			if sum == target {
				res++
			}
			return
		}
		// 遍历所有选择：+nums[i]、-nums[i]
		backtrack(i+1, sum+nums[i])
		backtrack(i+1, sum-nums[i])
	}
	backtrack(0, 0)
	return res
}

// ========================================
// 动态规划解法 + 详细注释
func findTargetSumWays(nums []int, target int) int {
	// 计算nums的和
	numsSum := 0
	for _, num := range nums {
		numsSum += num
	}

	// 计算negSum
	negSum := (numsSum - target) / 2

	// 判断边界条件
	if numsSum < target || (numsSum-target)%2 != 0 {
		// numsSum必须比target大
		// sum - target的值必须为偶数，不然(sum - target) / 2为小数，会不合题意（因为nums中都是正整数，其加和无法出现小数）
		// 直接返回0即可
		return 0
	}

	// 初始化dp表
	dp := make([][]int, len(nums)+1)
	for i := range dp {
		dp[i] = make([]int, negSum+1)
	}

	// 初始状态
    // 从前0个数字中选出若干个数字使得其和为0的方案数为1，即「空集合」不选任何数字即可得到0。
	dp[0][0] = 1

	// 状态转移
	for i := 1; i <= len(nums); i++ {
		for j := 0; j <= negSum; j++ {
			// 情况1：当前加和小于当前数字(j < nums[i-1])，我们不能放此数字
			//     => dp[i][j] = dp[i-1][j]
			// 情况2：当前加和大于等于当前数字(j >= nums[i-1])，我们可以选择放nums[i-1]也可以选择不放nums[i-1]，总方案数为其加和
			//     => dp[i][j] = dp[i-1][j] + dp[i-1][j-nums[i-1]]
			dp[i][j] = dp[i-1][j]
			if j >= nums[i-1] {
				dp[i][j] += dp[i-1][j-nums[i-1]]
			}
		}
	}

	return dp[len(nums)][negSum]
}

// @lc code=end

