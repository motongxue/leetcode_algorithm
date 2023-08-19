/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	// 按左端点进行排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{}

	// 采用pre、cur双指针
	pre := intervals[0]
	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		// 存在三种情况：
		if pre[1] >= cur[0] && cur[1] > pre[1] {
			// 1. pre和cur区间部分重叠
			pre[1] = cur[1]
		} else if cur[0] > pre[1] {
			// 2. pre和cur区间不重叠
			res = append(res, pre)
			pre = cur
		}
		// 3. pre和cur区间完全重叠，不需要处理
	}
	res = append(res, pre)
	return res
}

// @lc code=end

