/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */
package main

// @lc code=start
func lengthOfLongestSubstring(s string) int {
    windows := map[byte]int{}
    l, r := 0, 0
    cnt := 0
    for r < len(s) {
        ch := s[r]
        r++
        windows[ch]++
		// 收缩
        for windows[ch] > 1 {
            d := s[l]
            l++
            windows[d]--
        }
        // 此处不能直接使用len(mp)
        cnt = max(cnt, r-l)
    }
    return cnt
}
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

// @lc code=end
