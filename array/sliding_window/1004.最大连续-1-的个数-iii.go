package main

import (
	"fmt"
	"math"
)

/*
 * @lc app=leetcode.cn id=1004 lang=golang
 *
 * [1004] 最大连续1的个数 III
 */

// @lc code=start
func longestOnes(nums []int, k int) int {
	l, r, result := 0, 0, math.MinInt
	rule := k
	// 用来记录哪个下标是借用了k来替换
	mp := make(map[int]int)
	for r < len(nums) {
		// 满足条件
		// 如果还有k可以拿来替补，则还算满足条件
		if nums[r] != 1 && rule > 0 {
			mp[r] = 1
			rule--
		}
		// 不满足条件
		for l <= r && nums[r] == 0 && mp[r] == 0 && rule == 0 {
			if mp[l] == 1 {
				mp[l]--
				mp[r]++
			}
			l++
		}
		// 最大，则在外面更新
		result = max(result, r-l+1)
		r++
	}
	return result
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
func main() {
	fmt.Println(longestOnes([]int{0, 0, 1, 1, 1, 0, 0}, 0))
	fmt.Println(longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2))
	fmt.Println(longestOnes([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3))
}
