/*
 * @lc app=leetcode.cn id=1047 lang=golang
 *
 * [1047] 删除字符串中的所有相邻重复项
 */
package main
// @lc code=start
func removeDuplicates(s string) string {
	st := []rune{}
	for _,v := range s {
		if len(st) > 0 && v == st[len(st)-1] {
			st = st[:len(st)-1]
		}else {
			st = append(st, v)
		}
	}
	return string(st)
}
// @lc code=end

