/*
 * @lc app=leetcode.cn id=383 lang=golang
 *
 * [383] 赎金信
 */

// @lc code=start
func canConstruct(ransomNote string, magazine string) bool {
	mp := make(map[byte]int)
	for i := range magazine {
		mp[magazine[i]]++
	}
	for i := range ransomNote {
		if mp[ransomNote[i]] == 0 {
			return false
		}
		mp[ransomNote[i]]--
	}
	return true
}
// @lc code=end

