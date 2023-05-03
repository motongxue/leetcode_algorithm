/*
 * @lc app=leetcode.cn id=452 lang=golang
 *
 * [452] 用最少数量的箭引爆气球
 */

// @lc code=start
// 需要先进行排序
// 排序后，对于可射击的区间，有三种情况，情况3比较难想到
// 1. start1 start2 end1 end2
// 2. start1 end1 start2 end2
// 3. start1 start2 end2 end1
func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i,j int)bool {
		return points[i][0] < points[j][0]
	})
	end := points[0][1]
	res := 1
	for i := 1; i < len(points); i++ {
		if points[i][0] > end {
			end = points[i][1]
			res++
		}else if points[i][1] < end {
			// 此时需要缩小区间
			end = points[i][1]
		}
	}
	return res
}
// @lc code=end

