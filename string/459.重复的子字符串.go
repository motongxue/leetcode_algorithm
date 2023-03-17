/*
 * @lc app=leetcode.cn id=459 lang=golang
 *
 * [459] 重复的子字符串
 */
package main

import (
	"fmt"
)

// @lc code=start
func repeatedSubstringPattern(s string) bool {
	if len(s) == 1 {
		return false
	}
	ss := s + s
	ss = ss[1 : len(ss)-1]
	return find(ss, s)
}
func find(ss, s string) bool {
	next := make([]int, len(s))
	getNext(next, s)
	j := -1
	for i := 0; i < len(ss); i++ {
		for j >= 0 && ss[i] != s[j+1] {
			j = next[j]
		}
		if ss[i] == s[j+1] {
			j++
		}
		if j == len(s)-1 {
			return true
		}
	}
	return false
}
func getNext(next []int, s string) {
	j := -1
	next[0] = j
	for i := 1; i < len(s); i++ {
		for j >= 0 && s[i] != s[j+1] {
			j = next[j]
		}
		if s[i] == s[j+1] {
			j++
		}
		next[i] = j
	}
}

// @lc code=end
func main() {
	fmt.Println(repeatedSubstringPattern("abab"))
	fmt.Println(repeatedSubstringPattern("aba"))
	fmt.Println(repeatedSubstringPattern("abcabcabcabc"))
	fmt.Println(repeatedSubstringPattern("ss"))
}
