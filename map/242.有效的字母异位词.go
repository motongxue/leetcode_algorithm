/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	mp := make(map[byte]int)
	for i := range s {
		mp[s[i]]++
	}
	for i := range t {
		mp[t[i]]--
	}
	for _, v := range mp {
		if v != 0 {
			return false
		}
	}
	return true
}
// @lc code=end

