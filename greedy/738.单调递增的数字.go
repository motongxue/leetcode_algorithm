/*
 * @lc app=leetcode.cn id=738 lang=golang
 *
 * [738] 单调递增的数字
 */

// @lc code=start
func monotoneIncreasingDigits(N int) int {
	// 将int类型转为byte数组
    s := strconv.Itoa(N)
    ss := []byte(s)

    n := len(ss)
	// 记录需要修改的下标
    flag := n
    for i := n-1; i > 0; i-- {
        if ss[i] < ss[i-1] {
            flag = i
            ss[i-1]--
        }
    }
	// 将下标往后的所有数字修改为9（贪心）
    for i := flag; i < n; i++ {
        ss[i] = '9'
    }
	// 转为int类型
    res, _ := strconv.Atoi(string(ss))
    return res
}
// @lc code=end

