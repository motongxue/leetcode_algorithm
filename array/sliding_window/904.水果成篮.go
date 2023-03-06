/*
 * @lc app=leetcode.cn id=904 lang=golang
 *
 * [904] 水果成篮
 */
package main

// @lc code=start
func totalFruit(fruits []int) int {
	window := make(map[int]int)
	l, r := 0, 0
	result := 0
	for r < len(fruits) {
		c := fruits[r]
		r++
		window[c]++
		// 需要收缩
		for len(window) > 2 {
			d := fruits[l]
			l++		// 注意取完值后，再l++
			result = max(result, r-l)
			window[d]--
			if window[d] == 0 {
				delete(window, d)
			}
		}
	}
	// 可能由于达不到收缩条件，故需要单独讨论
	result = max(result, r-l)
	return result
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
