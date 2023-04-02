/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

package main

// @lc code=start
func letterCombinations(digits string) []string {
	res := []string{}
	if digits == "" {
		return res
	}
	mp := make(map[string]string)
	mp["2"] = "abc"
	mp["3"] = "def"
	mp["4"] = "ghi"
	mp["5"] = "jkl"
	mp["6"] = "mno"
	mp["7"] = "pqrs"
	mp["8"] = "tuv"
	mp["9"] = "wxyz"
	n := len(digits)
	var f func(i, n int, ss string)
	f = func(i, n int, ss string) {
		if i == n {
			if ss != "" {
				res = append(res, ss)
			}
			return
		}
		t := digits[i : i+1]
		for l := 0; l < len(mp[t]); l++ {
			f(i+1, n, ss+mp[t][l:l+1])
		}
	}
	f(0, n, "")
	return res
}

// @lc code=end
