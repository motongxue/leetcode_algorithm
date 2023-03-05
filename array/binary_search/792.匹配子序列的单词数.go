/*
 * @lc app=leetcode.cn id=792 lang=golang
 *
 * [792] 匹配子序列的单词数
 */

package main

// @lc code=start
func numMatchingSubseq(s string, words []string) int {
	// 初始化map
	pos := [26][]int{}
	for i := 0; i < len(s); i++ {
		pos[s[i]-'a'] = append(pos[s[i]-'a'], i)
	}
	ans := len(words)
	for _, w := range words {
		if len(words) > len(s) {
			ans--
			continue
		}
		j := 0
		for _, c := range w {
			if len(pos[c-'a']) == 0 {
				ans--
				break
			}
			idx := findL(pos[c-'a'], j)
			// idx := sort.SearchInts(pos[c-'a'], j)
			if idx == len(pos[c-'a']) {
				ans--
				break
			}
			j = pos[c-'a'][idx] + 1
		}

	}
	return ans
}

func findL(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] >= target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return r + 1
}

// @lc code=end
