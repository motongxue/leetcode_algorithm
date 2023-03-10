/*
 * @lc app=leetcode.cn id=202 lang=golang
 *
 * [202] 快乐数
 */

// @lc code=start
func isHappy(n int) bool {
	mp := make(map[int]bool)
	for n != 1 {
		tmp := 0
		for n != 0 {
			t := n % 10
			tmp += t * t
			n /= 10
		}
		n = tmp
		if mp[tmp] {
			return false
		}
		mp[tmp] = true
	}
	return true
}

// @lc code=end

