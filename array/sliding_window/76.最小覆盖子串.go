/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 */
package main

import (
	"math"
)

// @lc code=start
func minWindow(s string, t string) string {
	need := make(map[byte]int)   // 用于统计需要凑齐的字符
	window := make(map[byte]int) // 记录滑动窗口内已有字符的个数
	for i := range t {
		need[t[i]]++
	}
	left, right := 0, 0
	cnt := 0                          // 判断窗口中是否已经包含了字串 t 中所有字符
	start, length := 0, math.MaxInt32 // 最小覆盖子串的起始索引及长度
	for right < len(s) {
		c := s[right]
		right++
		// 如果这个字符在字串 t 中需要的话，则加入窗口
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] { // 如果字符 c 在窗口中的数量已经满足其在字串 t 中的数量
				cnt++ // 计数器加一
			}
		}
		// 当满足条件时，开始收缩
		for cnt == len(need) { // 如果滑动窗口中的字符已经完全覆盖字串 t 中的字符
			// 更新操作
			if right-left < length {
				start = left
				length = right - left
			}
			d := s[left] // d 是将要移出窗口的字符
			// 注意跟上面的对称性和顺序性
			left++                    // 左侧窗口右移
			if _, ok := need[d]; ok { // 如果这个字符在字串 t 中需要的话
				if window[d] == need[d] { // 如果这个字符已经满足了他在字串 t 中的需求
					cnt-- // 计数器减一
				}
				window[d]-- // 移出窗口
			}
		}
	}
	if length == math.MaxInt32 { // 如果最小子串长度没有更新，则返回空格
		return ""
	}
	return s[start : start+length] // 返回最小覆盖子串
}

// @lc code=end
