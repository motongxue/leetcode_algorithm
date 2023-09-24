/*
 * @lc app=leetcode.cn id=518 lang=golang
 *
 * [518] 零钱兑换 II
 */

// @lc code=start
// ====================== 递归 ======================
// 回溯解法：超时
func change(amount int, coins []int) int {
    res := 0
    n := len(coins)
    var dfs func(sum, step int)
    dfs = func(sum, step int) {
        if sum > amount {
            return
        }else if sum == amount {
            res++
            return
        }
        for i := step; i < n; i++ {
            dfs(sum+coins[i], i)
        }
    }
    dfs(0,0)
    return res
}
func change(amount int, coins []int) int {
    dp := make([]int, amount+1)
	// 必须对dp[0]进行初始化：初始化0大小的背包, 当然是不装任何东西了, 就是1种方法
    dp[0] = 1
	// 本题求的是组合数，所以外层循环是遍历物品，内层循环是遍历背包容量
    for i := 0; i < len(coins); i++ {			// 遍历物品
        for j := coins[i]; j <= amount; j++ {	// 遍历背包容量
            dp[j] += dp[j-coins[i]]
        }
    }
    return dp[amount]
}
// @lc code=end

