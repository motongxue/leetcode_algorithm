/*
 * @lc app=leetcode.cn id=367 lang=golang
 *
 * [367] 有效的完全平方数
 */

// @lc code=start
func isPerfectSquare(num int) bool {
	l, r := 0, 65535
	for l <= r {
		mid := (l + r) / 2
		if mid*mid > num {
			r = mid - 1
		} else if mid*mid < num {
			l = mid + 1
		} else {
			return true
		}
	}
	return false
}

// @lc code=end

