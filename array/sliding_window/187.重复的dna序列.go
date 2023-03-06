/*
 * @lc app=leetcode.cn id=187 lang=golang
 *
 * [187] 重复的DNA序列
 */

// @lc code=start
// 不像是滑动窗口，更像是for+map的操作
func findRepeatedDnaSequences(s string) []string {
	mp := make(map[string]int)
	result := []string{}
	i := 0
	for i+10 <= len(s) {
		if v, _ := mp[s[i:i+10]]; v == 1 {
			result = append(result, s[i:i+10])
		}
		mp[s[i:i+10]]++
		i++
	}
	return result
}

// @lc code=end

