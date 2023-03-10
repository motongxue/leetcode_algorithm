/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */
package main

import (
	"fmt"
	"sort"
)

// @lc code=start
//通过实现排序实现
func groupAnagrams(strs []string) [][]string {
    mp := map[string][]string{}
    for _, str := range strs {
        s := []byte(str)
		// TODO sort.Slice实现
        sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
        sortedStr := string(s)
        mp[sortedStr] = append(mp[sortedStr], str)
    }
    ans := make([][]string, 0, len(mp))
    for _, v := range mp {
        ans = append(ans, v)
    }
    return ans
}
// 此时key为一个数组，防止了相加后重复
// func groupAnagrams(strs []string) [][]string {
//     mp := map[[26]int][]string{}
//     for _, str := range strs {
//         cnt := [26]int{}
//         for _, b := range str {
//             cnt[b-'a']++
//         }
//         mp[cnt] = append(mp[cnt], str)
//     }
//     ans := make([][]string, 0, len(mp))
//     for _, v := range mp {
//         ans = append(ans, v)
//     }
//     return ans
// }

// @lc code=end
func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
