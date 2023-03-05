/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 */
package main

import (
	"fmt"
	"math"
)

// @lc code=start

// 最暴力的方法，超时
// func minSubArrayLen(target int, nums []int) int {
// 	res := -1
// 	step := len(nums) - 1
// 	for step >= 0 {
// 		i := 0
// 		j := i + step
// 		for i < len(nums) && j < len(nums) {
// 			tmp := 0
// 			for k := i; k <= j; k++ {
// 				tmp += nums[k]
// 			}
// 			if tmp >= target {
// 				res = step
// 			}
// 			i++
// 			j++
// 		}
// 		step--
// 	}
// 	if res == -1 {
// 		return 0
// 	} else {
// 		return res + 1
// 	}
// }

func minSubArrayLen(target int, nums []int) int {
	l, r, result := 0, 0, math.MaxInt
	sum := 0
	for r < len(nums) {
		sum += nums[r]
		for sum >= target {
			result = min(result, r-l+1)
			sum -= nums[l]
			l++
		}
		r++
	}
	if result == math.MaxInt {
		return 0
	}
	return result
	func min(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
}

// @lc code=end

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	fmt.Println(minSubArrayLen(4, []int{1, 4, 4}))
	fmt.Println(minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1}))
}
