/*
 * @lc app=leetcode.cn id=454 lang=golang
 *
 * [454] 四数相加 II
 */
package main

import "fmt"

// @lc code=start
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	mp := make(map[int]int)
	cnt := 0
	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			mp[v1+v2]++
		}
	}
	for _, v3 := range nums3 {
		for _, v4 := range nums4 {
			cnt += mp[-v3-v4]
		}
	}
	return cnt
}

// @lc code=end
func main() {
	fmt.Println(fourSumCount([]int{-1, -1}, []int{-1, 1}, []int{-1, 1}, []int{1, -1}))
}
