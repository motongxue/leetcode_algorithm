/*
 * @lc app=leetcode.cn id=219 lang=golang
 *
 * [219] 存在重复元素 II
 */

// @lc code=start

// 时间复杂度为O(n2)，超时
// func containsNearbyDuplicate(nums []int, k int) bool {
//     for i := 0; i < len(nums); i++ {
//         for j := i+1; j < len(nums); j++ {
//             if nums[i] == nums[j] && abs(i,j) <= k {
//                 return true
//             }
//         }
//     }
//     return false
// }
// func abs(a, b int) int {
//     if a > b {
//         return a - b
//     }
//     return b - a
// }
// 利用map，时间复杂度为O(n)
// func containsNearbyDuplicate(nums []int, k int) bool {
// 	利用map来降低复杂度
// 	mp := make(map[int][]int)
// 	for i := 0; i < len(nums); i++ {
// 		mp[nums[i]] = append(mp[nums[i]], i)
// 		if len(mp[nums[i]]) > 1 {
// 			// 存在
// 			// 由于是按顺序加进来，所以肯定相减后为正数
// 			if mp[nums[i]][1]-mp[nums[i]][0] <= k {
// 				return true
// 			} else {
// 				mp[nums[i]] = mp[nums[i]][1:]
// 			}
// 		}
// 	}
// 	return false
// }
// 利用滑动窗口
func containsNearbyDuplicate(nums []int, k int) bool {
	mp := make(map[int]bool)
	r := 0
	for r < len(nums) {
		if _, ok := mp[nums[r]]; ok {
			return true
		}
		mp[nums[r]] = true
		if len(mp) > k {
			delete(mp, nums[r-k])
		}
		r++
	}
	return false
}

// @lc code=end

