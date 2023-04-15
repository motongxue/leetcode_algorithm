/*
 * @lc app=leetcode.cn id=491 lang=golang
 *
 * [491] 递增子序列
 */

// @lc code=start
func findSubsequences(nums []int) [][]int {
	res := [][]int{}
	var backtracking func(startIdx int, path []int)
	backtracking = func(startIdx int, path []int) {
		if len(path) > 1 {
			tmp := []int{}
			tmp = append(tmp, path...)
			res = append(res, tmp)
		}
		used := make(map[int]bool, len(nums)) // 初始化used字典，用以对同层元素去重
		for i := startIdx; i < len(nums); i++ {
			// TODO 这种需要是有序情况下，才能保证同层不相同
			// if (i > startIdx && nums[i] == nums[i-1])
			if used[nums[i]] {
				continue
			}
			if len(path) == 0 || nums[i] >= path[len(path)-1] {
				used[nums[i]] = true
				backtracking(i+1, append(path, nums[i]))
			}
		}
	}
	backtracking(0, []int{})
	return res
}

// @lc code=end

