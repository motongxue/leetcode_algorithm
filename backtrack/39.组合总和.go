/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	var f func(target, start int, path []int)
	res := [][]int{}
	// 注意这里先排序了
	sort.Ints(candidates)
	f = func(target, start int, path []int) {
		if target < 0 {
			return
		} else if target == 0 {
			tmp := []int{}
			tmp = append(tmp, path...)
			res = append(res, tmp)
			return
		} else {
			for i := start; i < len(candidates); i++ {
				f(target-candidates[i], i, append(path, candidates[i]))
			}
		}
	}
	f(target, 0, []int{})
	return res
}

// @lc code=end

