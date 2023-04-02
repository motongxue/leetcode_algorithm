/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 */

// @lc code=start
func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}
	// 注意这里先排序了
	sort.Ints(candidates)
	var f func(target, start int, path []int)
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
				// TODO 注意，这里应为 i > start 而非 i > 0
				if i > start && candidates[i] == candidates[i-1] {
					continue
				}
				f(target-candidates[i], i+1, append(path, candidates[i]))
			}
		}
	}
	f(target, 0, []int{})
	return res
}

// @lc code=end

