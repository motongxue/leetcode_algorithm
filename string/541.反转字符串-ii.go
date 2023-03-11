/*
 * @lc app=leetcode.cn id=541 lang=golang
 *
 * [541] 反转字符串 II
 */
package main

// @lc code=start
func reverseStr(s string, k int) string {
	result := []byte(s)
	length := len(s)
	for i := 0; i < length; i += 2 * k {
		if i+k <= length {
			reverse(result, i, i+k-1)
		} else {
			reverse(result, i, length-1)
		}
	}
	return string(result)
}
func reverse(s []byte, i, j int) {
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

// @lc code=end
