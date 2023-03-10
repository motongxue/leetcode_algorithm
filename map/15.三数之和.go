/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */
package main

import "sort"

// @lc code=start
func threeSum(nums []int) [][]int {
	n := len(nums)
	result:=[][]int{}
	mp := make(map[[3]int]struct{})
	sort.Ints(nums)
	for i:=0;i<n;i++{
		for j:=i+1;j<n;j++{
			for k:=j+1;k<n;k++{
				if nums[i]+nums[j]+nums[k]==0{
					mp[[3]int{nums[i],nums[j],nums[k]}]=struct{}{}
				}
			}
		}
	}
	for k,_ := range mp {
		result = append(result,[]int{k[0],k[1],k[2]})
	}
	return result
}
// @lc code=end

