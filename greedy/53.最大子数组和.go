/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子数组和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	tmp, res := math.MinInt, math.MinInt
	for _, cur := range nums {
		// tmp < 0, cur < 0	    -> tmp = max(tmp, cur)
		// 			cur = 0     -> tmp = cur
		// 			cur > 0     -> tmp = cur
		// tmp > 0, cur >,<,= 0 -> tmp += cur 
		if tmp < 0 && cur < 0 {
			tmp = max(tmp, cur)
		}else if tmp < 0 && cur >= 0 {
			// 注意这个地方需要考虑cur=0的情况
			tmp = cur
		}else {
			tmp += cur
		}
		res = max(res, tmp)
	}
	return res
}
func max(a,b int) int {
	if a < b {
		return b
	}
	return a
}
// @lc code=end

