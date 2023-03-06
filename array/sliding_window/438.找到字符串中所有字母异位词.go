/*
 * @lc app=leetcode.cn id=438 lang=golang
 *
 * [438] 找到字符串中所有字母异位词
 */
package main

import "fmt"

// @lc code=start
func findAnagrams(s string, p string) []int {
	// s 是长的，p 是短的
	need := make(map[byte]int)
	window := make(map[byte]int)
	l, r, cnt := 0, 0, 0
	result := []int{}
	for i := range p {
		need[p[i]]++
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
		// 需要收缩时
		for r-l >= len(p) {
			if cnt == len(need) {
				result = append(result, l)
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
	return result
}

// @lc code=end
func main() {
	fmt.Println(findAnagrams("abaacbabc", "abc"))
}
