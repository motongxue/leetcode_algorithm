/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 */

// @lc code=start
func mySqrt(x int) int {
	l, r := 0, 65535
	for l <= r {
		mid := (l + r) / 2
		if mid*mid <= x {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l - 1
}

// @lc code=end

