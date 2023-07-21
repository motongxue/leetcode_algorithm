/*
 * @lc app=leetcode.cn id=474 lang=golang
 *
 * [474] 一和零
 */

// @lc code=start
func findMaxForm(strs []string, m int, n int) int {
    // 一个背包，但有两个限制，需要同时满足限制才能加入背包，则采用二维数组。
    // dp[i][j]：最多有i个0和j个1的strs的最大子集的大小为dp[i][j]。
    dp := make([][]int, m+1)
    for i := 0; i <= m; i++ {
        dp[i] = make([]int, n+1)
    }
    // 相当于在遍历物品
    for _, ss := range strs {
        zero, one := cal(ss)
        for i := m; i >= zero; i-- {
            for j := n; j >= one; j-- {
                // 不需要考虑字符串的长度问题，因为此时的dp含义是最大子集的大小，通过max比较后，自然会取最大的
                dp[i][j] = max(dp[i][j], dp[i-zero][j-one] + 1)
            }
        }
    }

    return dp[m][n]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

// 定义一个函数，传入字符串，返回有多少个0和1
func cal(ss string) (zero, one int) {
    for _, v := range ss {
        if v == '0' {
            zero++
        }else {
            one++
        }
    }
    return
}
// @lc code=end