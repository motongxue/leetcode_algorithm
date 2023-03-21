/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */
package main

// @lc code=start
func isValid(s string) bool {
	st := []rune{}
	for _, v := range s {
		if v == '(' {
			st = append(st, ')')
		} else if v == '{' {
			st = append(st, '}')
		} else if v == '[' {
			st = append(st, ']')
		} else if len(st) > 0 && st[len(st)-1] == v {
			st = st[:len(st)-1]
		} else {
			return false
		}
	}
	if len(st) == 0 {
		return true
	} else {
		return false
	}
}

// @lc code=end
