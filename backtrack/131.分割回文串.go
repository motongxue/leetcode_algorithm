/*
 * @lc app=leetcode.cn id=131 lang=golang
 *
 * [131] 分割回文串
 */
package main
// @lc code=start
// ======================================
// 这里的回文串判断可以优化
// 几个难点：
// 切割问题可以抽象为组合问题
// 如何模拟那些切割线
// 切割问题中递归如何终止
// 在递归循环中如何截取子串
// 如何判断回文
func partition(s string) [][]string {
	res := [][]string{}
	var backtracking func(s string, path []string, startIdx int)
	backtracking = func(s string, path []string, startIdx int) {
		if startIdx >= len(s) {
			tmp := []string{}
			tmp = append(tmp, path...)
			res = append(res, tmp)
			return
		}
		for i := startIdx; i < len(s); i++ {
			if isPalindrome(s,startIdx,i) {
				backtracking(s,append(path,s[startIdx:i+1]),i+1)
			}
		}
	}
	backtracking(s,[]string{},0)
	return res
}
func isPalindrome(s string, l, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
// @lc code=end

