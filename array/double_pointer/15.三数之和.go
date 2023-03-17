/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */
package main

import (
	"sort"
)

// @lc code=start
// i从下标0的地方开始，同时定一个下标left 定义在i+1的位置上，
// 定义下标right 在数组结尾的位置上。
// ->所以实际上是 i, left, right 这种位置顺序
func threeSum(nums []int) [][]int {
	n := len(nums)
	result := [][]int{}
	sort.Ints(nums)
	for i := 0; i < n; i++ {
		//  排序之后如果第一个元素已经大于零
		if nums[i] > 0 {
			return result
		}
		// 对nums[i]进行去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, n-1
		for left < right {
			x := nums[left] + nums[right] + nums[i]
			if x == 0 {
				tmp := []int{nums[left], nums[i], nums[right]}
				// 由于本身有位置顺序，并且已经去掉重复元素，故不包含重复的三元组
				result = append(result, tmp)
				// 对left和right所指向元素去重
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				// 指向新的元素
				right--
				left++
			} else if x > 0 {
				// 如果nums[i] + nums[left] + nums[right] > 0
				// right下标就应该向左移动
				right--
			} else {
				// 应向右移动
				left++
			}
		}
	}
	return result
}

// @lc code=end
