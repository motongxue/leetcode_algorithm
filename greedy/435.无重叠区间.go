/*
 * @lc app=leetcode.cn id=435 lang=golang
 *
 * [435] 无重叠区间
 */

// @lc code=start
// ============================================
// 解法一
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

// ============================================
// 解法二：同样利用重叠子集的相对问题来求解
// 但注意这里的计数时，端点可以重叠
// 只要把弓箭那道题目代码里射爆气球的判断条件加个等号（认为[0，1][1，2]不是相邻区间）
func eraseOverlapIntervals(intervals [][]int) int {
	cnt := 0
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	end := math.MinInt
	for _, v := range intervals {
		if v[0] >= end {
			cnt++
			end = v[1]
		} else if v[1] < end {
			end = v[1]
		}
	}
	return len(intervals) - cnt
}

// @lc code=end

