/*
 * @lc app=leetcode.cn id=93 lang=golang
 *
 * [93] 复原 IP 地址
 */
package main

import (
	"strconv"
	"strings"
)

// @lc code=start
func restoreIpAddresses(s string) []string {
	res := []string{}
	var backtracking func(startIdx int, path string)
	backtracking = func(startIdx int, path string) {
		// 由于每次都 +"."，所以Split拆分后需要有5个值（最后一个为0）
		if startIdx >= len(s) && len(strings.Split(path, ".")) == 5 {
			tmp := path[:len(path)-1]
			res = append(res, tmp)
			return
		}
		for i := startIdx; i < len(s); i++ {
			t, _ := strconv.Atoi(s[startIdx : i+1])
			// t需要小于255，数值>255的表示该ip已失效
			// 当有前导0时也失效
			if t > 255 || (i > startIdx && s[startIdx:i] == "0") {
				return
			}
			backtracking(i+1, path+s[startIdx:i+1]+".")
		}
	}
	backtracking(0, "")
	return res
}

// @lc code=end
