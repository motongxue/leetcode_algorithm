/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] 两个数组的交集
 */

package main

import (
	"fmt"
)

// @lc code=start

func intersection(nums1 []int, nums2 []int) []int {
	n := len(nums1) - 1
	result := []int{}
	set := make(map[int]int, 0)
	for _, v := range nums2 {
		l, r := 0, n
		for l <= r {
			mid := (l + r) / 2
			if nums1[mid] > v {
				r = mid - 1
			} else if nums1[mid] < v {
				l = mid + 1
			} else {
				set[v] = 1
				break
			}
		}
	}
	for k, _ := range set {
		result = append(result, k)
	}
	return result
}

// @lc code=end
func main() {
	fmt.Println(intersection([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}
