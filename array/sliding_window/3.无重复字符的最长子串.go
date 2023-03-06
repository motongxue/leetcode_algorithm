/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */
package main

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	l, r := 0, 0
	mp := make(map[byte]int)
	cnt := 0
	for r < len(s) {
		c := s[r]
		mp[c]++
		r++
		for mp[c] > 1 {
			d := s[l]
			mp[d]--
			l++
		}
		cnt = max(r-l, cnt)
	}
	return cnt
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
