/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 四数之和
 */
package main

import (
	"fmt"
	"sort"
)

// @lc code=start
func fourSum(nums []int, target int) [][]int {
	n := len(nums)
	result := [][]int{}
	if n < 4 {
		return nil
	}
	sort.Ints(nums)
	for i1 := 0; i1 < n-3; i1++ {
		if i1 > 0 && nums[i1] == nums[i1-1] {
			continue
		}
		// 本质上同三数之和原理相同，甚至可以类比到 n数之和
		// 先找到i1，然后令i2为i1+1，此后，再找i3, i4 的位置
		for i2 := i1 + 1; i2 < n-2; i2++ {
			if i2 > i1+1 && nums[i2] == nums[i2-1] {
				continue
			}
			i3, i4 := i2+1, n-1
			for i3 < i4 {
				x := nums[i1] + nums[i2] + nums[i3] + nums[i4]
				if x == target {
					result = append(result, []int{nums[i1], nums[i2], nums[i3], nums[i4]})
					for i3 < i4 && nums[i4] == nums[i4-1] {
						i4--
					}
					for i3 < i4 && nums[i3] == nums[i3+1] {
						i3++
					}
					i3++
					i4--
				} else if x > target {
					i4--
				} else {
					i3++
				}
			}
		}
	}
	return result
}

// @lc code=end

func main() {
	fmt.Println(fourSum([]int{1, -2, -5, -4, -3, 3, 3, 5}, -11)) // [[-5,-4,-3,1]]
	fmt.Println(fourSum([]int{-3, -1, 0, 2, 4, 5}, 0))
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	fmt.Println(fourSum([]int{2, 2, 2, 2, 2}, 8))
	// dfs(0, 5)
}
