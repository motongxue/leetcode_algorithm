/*
 * @lc app=leetcode.cn id=216 lang=golang
 *
 * [216] 组合总和 III
 */
package main

// @lc code=start
func combinationSum3(k int, n int) [][]int {
	var f func(k, n, start int, path []int)
	res := [][]int{}
	f = func(k, n, start int, path []int) {
		if k == 0 {
			if n == 0 {
				tmp := []int{}
				tmp = append(tmp, path...)
				res = append(res, tmp)
			}
			return
		}
		for start < 10 {
			if n-start >= 0 {
				f(k-1, n-start, start+1, append(path, start))
			}
			start++
		}
	}
	f(k, n, 1, []int{})
	return res
}
// @lc code=end
