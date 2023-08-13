/*
 * @lc app=leetcode.cn id=452 lang=golang
 *
 * [452] 用最少数量的箭引爆气球
 */

// @lc code=start
// ============================================
// 解法一：利用区间进行记录，比较直观，但该题不需要记录区间起点
func findMinArrowShots(points [][]int) int {
	res := 0
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	curP := make([]int, 2)
	curP[0], curP[1] = math.MinInt, math.MinInt
	for i := 0; i < len(points); i++ {
		newP := points[i]
		// 如果新的点的起点已经比当前的终点还大
		if newP[0] > curP[1] {
			res++
			curP = newP
		} else {
			// 记录新的起点和终点
			curP[0] = newP[0]
			curP[1] = min(curP[1], newP[1])
		}
	}
	return res
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// ============================================
// 解法二：
// 需要先进行排序
// 排序后，对于可射击的区间，有三种情况，情况3比较难想到
// 1. start1 start2 end1 end2
// 2. start1 end1 start2 end2
// 3. start1 start2 end2 end1
func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	end := points[0][1]
	res := 1
	for i := 1; i < len(points); i++ {
		if points[i][0] > end {
			end = points[i][1]
			res++
		} else if points[i][1] < end {
			// 此时需要缩小区间
			end = points[i][1]
		}
	}
	return res
}

// @lc code=end
