/*
 * @lc app=leetcode.cn id=567 lang=golang
 *
 * [567] 字符串的排列
 */
package main

import "fmt"

// @lc code=start
func checkInclusion(t string, s string) bool {
	need := make(map[byte]int)
	window := make(map[byte]int)
	l, r := 0, 0
	cnt := 0
	for i := range t {
		need[t[i]]++
	}
	for r < len(s) {
		c := s[r]
		r++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				cnt++
			}
		}
		// fmt.Printf("window: [%d, %d)\n", l, r)
		for r-l >= len(t) {
			if cnt == len(need) {
				return true
			}
			d := s[l]
			l++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					cnt--
				}
				window[d]--
			}
		}
	}
	return false
}

// @lc code=end

func main() {
	fmt.Println(checkInclusion("abcdxabcde", "abcdeabcdx"))
}
