/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */
package main

import "fmt"

// @lc code=start
// func combine(n int, k int) [][]int {
// 	res := [][]int{}
// 	// TODO 这一步不可省略，否则会提示函数f undefined
// 	var f func(n int, k int, path []int)
// 	f = func(n int, k int, path []int) {
// 		if k == 0 {
// 			tmp := []int{}
// 			tmp = append(tmp, path...)
// 			res = append(res, tmp)
// 			return
// 		}
// 		for i := n; i > 0; i-- {
// 			f(i-1, k-1, append(path, i))
// 		}
// 	}
// 	f(n, k, []int{})
// 	return res
// }
func combine(n int, k int) [][]int {
	res := [][]int{}
	var f func(startIdx int, path []int)
	f = func(startIdx int, path []int) {
		if len(path) == k {
			tmp := make([]int, k)
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		for i := startIdx; i <= n; i++ {
			f(i+1, append(path, i))
		}
	}
	f(1, []int{})
	return res
}

// @lc code=end
func main() {
	fmt.Println(combine(4, 2))
}
