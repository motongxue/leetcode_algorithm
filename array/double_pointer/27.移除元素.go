/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */
package main

import "fmt"

// @lc code=start
func removeElement(nums []int, val int) int {
	r := len(nums) - 1

	for l := 0; l <= r; l++ {
		for l <= r && nums[r] == val {
			r--
		}
		if l <= r && nums[l] == val {
			nums[l] = nums[r]
			r--
		}
	}
	return r + 1
}
// @lc code=end

func main() {
	fmt.Println(removeElement([]int{5}, 5))
	fmt.Println(removeElement([]int{4, 5}, 5))
	fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))
	fmt.Println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}
