/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 */

// @lc code=start
func maxProfit(prices []int) int {
	res := 0
	if len(prices) == 1 {
		return res
	}
	// 核心是需要求出上升的数量
	for i := 1; i < len(prices); i++ {
		if prices[i] - prices[i-1] >= 0{
			res += prices[i] - prices[i-1]
		}
	}
	return res
}
// @lc code=end

