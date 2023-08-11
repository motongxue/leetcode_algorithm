/*
 * @lc app=leetcode.cn id=135 lang=golang
 *
 * [135] 分发糖果
 */

// @lc code=start
// 将「相邻的孩子中，评分高的孩子必须获得更多的糖果」这句话拆分为两个规则，分别处理。
// 左规则：当 ratings[i−1]<ratings[i]时，i号学生的糖果数量将比i−1号孩子的糖果数量多。
// 右规则：当 ratings[i]>ratings[i+1]时，i号学生的糖果数量将比i+1号孩子的糖果数量多。
func candy(ratings []int) int {
	left := make([]int, len(ratings))
	right := make([]int, len(ratings))
	for i, _ := range ratings {
		left[i] = 1
		right[i] = 1
	}
	// 使满足左规则
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		}
	}
	// 使满足右规则
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			right[i] = right[i+1] + 1
		}
	}
	res := 0
	// 取以上2轮遍历left和right对应学生糖果数的最大值，
	// 这样则同时满足左规则和右规则，即得到每个同学的最少糖果数量。
	for i, _ := range left {
		res += max(left[i], right[i])
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end