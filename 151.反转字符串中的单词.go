/*
 * @lc app=leetcode.cn id=151 lang=golang
 *
 * [151] 反转字符串中的单词
 */
package main

import (
	"fmt"
	"strings"
)

// @lc code=start
func reverseWords(s string) string {
	// TODO 分割多个空格，则用Fields
	ss := strings.Fields(s)
	result := []byte{}
	length := len(ss) - 1
	for length > 0 {
		result = append(result, []byte(ss[length]+" ")...)
		length--
	}
	result = append(result, []byte(ss[length])...)
	return string(result)
}

// @lc code=end
func main() {
	fmt.Println(reverseWords("  hello world  "))
}
