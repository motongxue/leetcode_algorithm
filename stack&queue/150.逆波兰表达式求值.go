/*
 * @lc app=leetcode.cn id=150 lang=golang
 *
 * [150] 逆波兰表达式求值
 */
package main

import "strconv"

// @lc code=start
func evalRPN(tokens []string) int {
	st := []int{}
	for _, v := range tokens {
		if v == "+" {
			b := st[len(st)-2]
			a := st[len(st)-1]
			st = st[:len(st)-2]
			st = append(st, a+b)
		} else if v == "-" {
			b := st[len(st)-2]
			a := st[len(st)-1]
			st = st[:len(st)-2]
			st = append(st, b-a)
		} else if v == "*" {
			b := st[len(st)-2]
			a := st[len(st)-1]
			st = st[:len(st)-2]
			st = append(st, b*a)
		} else if v == "/" {
			b := st[len(st)-2]
			a := st[len(st)-1]
			st = st[:len(st)-2]
			st = append(st, b/a)
		} else {
			v, _ := strconv.Atoi(v)
			st = append(st, v)
		}
	}
	return st[0]
}

// @lc code=end
