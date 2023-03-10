/*
 * @lc app=leetcode.cn id=350 lang=golang
 *
 * [350] 两个数组的交集 II
 */

// @lc code=start
func intersect(nums1 []int, nums2 []int) []int {
	mp := make(map[int]int)
	for _, v := range nums1 {
		mp[v]++
	}
	result := []int{}
	for _, v := range nums2 {
		if _, ok := mp[v]; ok {
			result = append(result, v)
			mp[v]--
			if mp[v] == 0 {
				delete(mp, v)
			}
		}
	}
	return result
}

// @lc code=end

