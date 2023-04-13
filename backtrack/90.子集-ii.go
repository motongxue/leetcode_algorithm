/*
 * @lc app=leetcode.cn id=90 lang=golang
 *
 * [90] 子集 II
 */
package main

import (
	"fmt"
	"sort"
)

// @lc code=start
// ========================
// 版本一：使用set进行去重
//
//	func subsetsWithDup(nums []int) [][]int {
//		res := [][]int{}
//		set := make(map[[10]int]bool)
//
// 排序是必须的
//
//		sort.Ints(nums)
//		var backtracking func(startIdx int, path []int)
//		backtracking = func(startIdx int, path []int) {
//			// 难的是对切片进行去重
//			tmp := [10]int{}
//			i := 0
//			for i < len(path) {
//				tmp[i] = path[i]
//				i++
//			}
//			for i < 10 {
//				// 之所以要多这一步是因为默认空集的tmp为全0，跟元素为0的集合一样，导致0添加不进
//				tmp[i] = 99
//				i++
//			}
//			// 去重结束
//			if !set[tmp] {
//				t := []int{}
//				t = append(t, path...)
//				res = append(res, t)
//				set[tmp] = true
//			}
//			for i := startIdx; i < len(nums); i++ {
//				backtracking(i+1, append(path, nums[i]))
//			}
//		}
//		backtracking(0, []int{})
//		return res
//	}
//
// ================================================
// 不适用set版本：在递归前便进行筛选
func subsetsWithDup(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	var backtracking func(startIdx int, path []int)
	backtracking = func(startIdx int, path []int) {
		tmp := []int{}
		tmp = append(tmp, path...)
		res = append(res, tmp)

		for i := startIdx; i < len(nums); i++ {
			// TODO 同样注意这里的 i > startIdx，
			// 并且nums[i] == nums[i-1]指的是对同一树层使用过的元素进行跳过
			if i > startIdx && nums[i] == nums[i-1] {
				continue
			}
			backtracking(i+1, append(path, nums[i]))
		}
	}
	backtracking(0, []int{})
	return res
}

// @lc code=end

func main() {
	fmt.Println(subsetsWithDup([]int{0}))
}
