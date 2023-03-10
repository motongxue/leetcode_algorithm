/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] 两个数组的交集
 */
package main

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	mp := make(map[int]struct{})
	for _, v := range nums1 {
		mp[v] = struct{}{}
	}
	result := []int{}
	for _, v := range nums2 {
		if _, ok := mp[v]; ok {
			result = append(result, v)
			delete(mp, v)
		}
	}
	return result
}

// @lc code=end
