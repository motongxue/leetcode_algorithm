/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start
func subsets(nums []int) [][]int {
	res := [][]int{}
	var backtracking func(startIdx int, path []int)
	backtracking = func(startIdx int, path []int) {
		// 因为每一次的值都是符合要求的值，所以每次进来都进行存值，由于在这，刚好把[]空集处理了
		tmp := []int{}
		tmp = append(tmp, path...)
		res = append(res, tmp)
		// if startIdx == len(nums) {
		// 	return
		// }
		// 若 startIdx == len(nums) 时，也不会进入for循环，所以可以取消
		for i := startIdx; i < len(nums); i++ {
			backtracking(i+1, append(path, nums[i]))
		}
	}
	backtracking(0, []int{})
	return res
}

// @lc code=end

