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
func minSubArrayLen(target int, nums []int) int {
	// 此时窗口变成了整型
    window := 0
    l, r := 0, 0
    res := math.MaxInt
    for r < len(nums) {
        c := nums[r]
        r++
        window += c

		// 当窗口中的和已经满足≥target的条件后，便可以进行收缩
        for window >= target {
			// 需要先进行统计
            res = min(res, r - l)
            d := nums[l]
            l++
            window -= d
        }
    }
    if res == math.MaxInt {
        return 0
    }
    return res
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

// @lc code=end

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	fmt.Println(minSubArrayLen(4, []int{1, 4, 4}))
	fmt.Println(minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1}))
}
