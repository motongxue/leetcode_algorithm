/*
 * @lc app=leetcode.cn id=435 lang=golang
 *
 * [435] 无重叠区间
 */

// @lc code=start
// 可以按照区间左端点排序，只要从后向前遍历即可
func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := 1
	start := intervals[len(intervals)-1][0]
	for i := len(intervals) - 2; i >= 0; i-- {
		if intervals[i][1] <= start {
			res++
			start = intervals[i][0]
		}
	}
	return len(intervals) - res
}

// @lc code=end

