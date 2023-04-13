/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
func permute(nums []int) [][]int {
	res := [][]int{}
	vis := make(map[int]bool)
	var backtrace func(path []int)
	backtrace = func(path []int) {
		if len(path) == len(nums) {
			tmp := []int{}
			tmp = append(tmp, path...)
			res = append(res, tmp)
		}
		for i := 0; i < len(nums); i++ {
			if !vis[nums[i]] {
				vis[nums[i]] = true
				backtrace(append(path, nums[i]))
				vis[nums[i]] = false
			}
		}
	}
	backtrace([]int{})
	return res
}

// @lc code=end

