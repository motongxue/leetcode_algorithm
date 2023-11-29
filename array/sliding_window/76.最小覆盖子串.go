/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 */
package main

import (
	"math"
)

// @lc code=start
func minWindow(s string, t string) string {
    // need用来记录t中的字符数量
    need := make(map[byte]int)
    for i := range t {
		need[t[i]]++
	}
    // 记录此时窗口中已经拥有的字符
    window := make(map[byte]int)
    // 滑动窗口左右指针
    l, r := 0, 0
    // cnt：判断窗口中是否已经包含了字串t中所有字符
    // res：长度判断
    cnt, res := 0, math.MaxInt
    // sl、sr：返回字符串的下标
    sl, sr := 0, len(s)-1
    for r < len(s) {
        c := s[r]
        r++

        window[c]++
        if window[c] == need[c] { // 如果字符 c 在窗口中的数量已经满足其在字串 t 中的数量
            cnt++ // 计数器加一
        }

        // 如果window已经包含了所有的t，则进行收缩，并保证最小
        for cnt == len(need) {
            d := s[l]
            l++

            // 当此时d为need中有的字符时，则需要统计
            if window[d] == need[d] {
                cnt--
                if res > r - l {
                    res = r - l
                    sl, sr = l-1, r
                }
            }
            window[d]--
        }
    }
    if res == math.MaxInt {
        return ""
    }
    return s[sl:sr]
}

// @lc code=end
