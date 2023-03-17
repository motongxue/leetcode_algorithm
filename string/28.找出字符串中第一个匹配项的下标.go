/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 找出字符串中第一个匹配项的下标
 */
package main

// @lc code=start
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	// TODO 这种写法可以定义大小一定的[]int
	next := make([]int, len(needle))
	getNext(next, needle)
	// 注意与next中的书写的差别
	// i表示前缀末尾、j表示后缀末尾
	j := -1
	// 注意此时对haystack进行遍历
	for i := 0; i < len(haystack); i++ {
		for j >= 0 && haystack[i] != needle[j+1] {
			j = next[j]
		}
		if haystack[i] == needle[j+1] {
			j++
		}
		// 如果已经到达末尾，则返回
		if j == len(needle)-1 {
			return i - len(needle) + 1
		}
	}
	return -1
}
func getNext(next []int, needle string) {
	j := -1
	next[0] = j
	// 注意此时对needle进行遍历
	for i := 1; i < len(needle); i++ {
		// 这里注意是 >=0 还有 s[j+1]
		// 向前回退可能是一个连续的过程，所以需要用for
		for j >= 0 && needle[i] != needle[j+1] {
			j = next[j]
		}
		if needle[i] == needle[j+1] {
			j++
		}
		// 完成后对next[i]进行赋值
		next[i] = j
	}
}

// @lc code=end
