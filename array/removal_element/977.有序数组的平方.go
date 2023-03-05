/*
 * @lc app=leetcode.cn id=977 lang=golang
 *
 * [977] 有序数组的平方
 */
package main

import "fmt"

// @lc code=start
func sortedSquares(nums []int) []int {
	n := len(nums)
	i, j, k := 0, n-1, n-1
	ans := make([]int, n)
	for i <= j {
		if nums[i]*nums[i] < nums[j]*nums[j] {
			ans[k] = nums[j] * nums[j]
			j--
		} else {
			ans[k] = nums[i] * nums[i]
			i++
		}
		k--
	}
	return ans
}

// func sortedSquares(nums []int) []int {
//     ans := []int{}
//     for range nums {
//         index := binarySearch(nums, 0)
//         ans = append(ans, nums[index] * nums[index])
//         nums = append(nums[:index], nums[index+1:]...)
//     }
//     return ans
// }

// func binarySearch(nums []int, target int) int {
//     i := search(len(nums), func(i int) bool {return nums[i] >= target})
//     if i < len(nums) && nums[i] == 0 {
//         return i
//     }
//     // 如果没找到，但是 靠近0 的数字的索引在 0 < i < len(nums)
//     if 0 < i && i < len(nums) {// 如果有正数这个索引肯定落在 正数区域
//         if 0 - nums[i-1] < nums[i] {
//             return i-1
//         }
//         return i
//     }
//     // 处理边界问题
//     if 0 == i {
//         return i
//     }
//     return i-1
// }

// func search(n int, f func(int) bool) int {
//     i, j := 0, n
//     for i < j {
//         h := int(uint(i+j)>>1)
//         if !f(h) {
//             i = h + 1
//         }else {
//             j = h
//         }
//     }
//     return i
// }

// @lc code=end

func main() {
	fmt.Println(sortedSquares([]int{0}))
	fmt.Println(sortedSquares([]int{1, 1}))
	fmt.Println(sortedSquares([]int{1, 2}))
	fmt.Println(sortedSquares([]int{-4, 2}))
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
	fmt.Println(sortedSquares([]int{-7, -3, 2, 3, 11}))
	fmt.Println(sortedSquares([]int{-5, -3, -2, -1}))
	fmt.Println(sortedSquares([]int{-1, 2, 2}))
}
