package main

import "fmt"

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			return mid
		}
	}
	return r + 1
}
func main() {
	fmt.Println(searchInsert([]int{0, 1, 3}, 2))
}
