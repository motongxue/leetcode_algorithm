/*
 * @lc app=leetcode.cn id=844 lang=golang
 *
 * [844] 比较含退格的字符串
 */
package main

import "fmt"

// @lc code=start
func backspaceCompare(s string, t string) bool {
	i1 := 0
	r1, r2 := make([]byte, len(s)), make([]byte, len(t))
	for j1 := 0; j1 < len(s); j1++ {
		if s[j1] != '#' {
			r1[i1] = s[j1]
			i1++
		} else {
			if i1 == 0 {
				continue
			}
			i1--
		}
	}
	i2 := 0
	for j2 := 0; j2 < len(t); j2++ {
		if t[j2] != '#' {
			r2[i2] = t[j2]
			i2++
		} else {

			if i2 == 0 {
				continue
			}
			i2--
		}
	}
	if i1 != i2 {
		return false
	}
	for i := 0; i < i1; i++ {
		if r1[i] != r2[i] {
			return false
		}
	}
	return true
}

// @lc code=end
func main() {
	fmt.Println(backspaceCompare("ab#c", "ab#c"))
}
