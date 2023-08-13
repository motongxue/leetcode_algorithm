/*
 * @lc app=leetcode.cn id=763 lang=golang
 *
 * [763] 划分字母区间
 */

// @lc code=start
// 思路：
// 利用idx统计每个字符的最远距离，
// 由于同一个字母只能出现在同一个片段，
// 显然同一个字母的第一次出现的下标位置和最后一次出现的下标位置必须出现在同一个片段。
// 并且利用更新right，从而达到最大包含
func partitionLabels(s string) []int {
	res := []int{}
	arr := make([]int, 26)
	for i := 0; i < len(s); i++ {
		arr[s[i]-'a'] = i
	}
	maxIdx := 0
	start := -1
	for i := 0; i < len(s); i++ {
		// 当已经到达距离开头最远的点了，那么此时其arr为本身！
		maxIdx = max(maxIdx, arr[s[i]-'a'])
		if maxIdx == i {
			res = append(res, i-start)
			start = i
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

