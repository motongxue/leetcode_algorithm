/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */
package main

import (
	"fmt"
	"math"
)

// @lc code=start
var (
	mp map[int]int
)

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	l, r, result := 0, 0, math.MinInt
	mp = make(map[int]int)
	for r < len(s) {
		// 不满足条件
		mp[int(s[r])]++
		// while 满足条件
		for check() {
			mp[int(s[l])]--
			l++
		}
		result = max(result, r-l+1)
		r++
	}
	return result
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 可能是右指针元素重复了，所以不能仅判断mp[l]
func check() bool {
	flag := false
	for _, v := range mp {
		if v > 1 {
			flag = true
		}
	}
	return flag
}

// @lc code=end

func main() {
	// fmt.Println(lengthOfLongestSubstring(""))
	// fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
