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
		// 总利润由每天的利润构成
		// 如果有钱赚，则此刻赚入
		if prices[i] - prices[i-1] >= 0{
			res += prices[i] - prices[i-1]
		}
	}
	return res
}
// 思路二：定义买入和卖出时机
func maxProfit(prices []int) int {
    res := 0
    if len(prices) == 1 {
        return 0
    }
    buy, sell := prices[0], prices[0]
    for i := 1; i < len(prices); i++ {
		// 如果价格变低，则将之前赚的记录下来，然后再从此刻买入
        if prices[i] - prices[i-1] < 0 {
            res += sell - buy
            buy = prices[i]
        }
        sell = prices[i]
    }
    if buy != sell {
        res += sell - buy
    }
    return res
}
// @lc code=end

