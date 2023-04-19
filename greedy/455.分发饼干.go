/*
 * @lc app=leetcode.cn id=455 lang=golang
 *
 * [455] 分发饼干
 */

// @lc code=start
func findContentChildren(g []int, s []int) int {
    sort.Ints(g)
    sort.Ints(s)
    res := 0
    i, j := 0, 0
    for i < len(g) && j < len(s) {
        if s[j] >= g[i] {
            i++
            res++
        }
        j++
    }
    return res
}
// @lc code=end

