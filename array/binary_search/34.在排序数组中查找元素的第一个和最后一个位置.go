/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	L, R := findLeft(nums, target), findRight(nums, target)
	if L > R {
		L, R = -1, -1
	}
	return []int{L, R}
}
func findRight(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] > target {
			r = mid - 1
			} else {
				// nums[mid] <= target
				l = mid + 1
			}
		}
		return l - 1
	}
	func findLeft(nums []int, target int) int {
		l, r := 0, len(nums)-1
		for l <= r {
			mid := (l + r) / 2
			if nums[mid] >= target {
				r = mid - 1
			} else {
			// nums[mid] <= target
			l = mid + 1
		}
	}
	return r + 1
}
// @lc code=end

