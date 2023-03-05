/*
 * @lc app=leetcode.cn id=658 lang=golang
 *
 * [658] 找到 K 个最接近的元素
 */

 package main
 
 import "fmt"
// @lc code=start

func findClosestElements(arr []int, k int, x int) []int {
	l, r := 0, len(arr)-1
	for l <= r {
		mid := (l + r) / 2
		if arr[mid] >= x {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	l = r + 1
	r = l + 1
	if l == len(arr) || arr[l] > x {
		l--
		r--
	}
	for k != 0 {
		k--
		if l == -1 {
			r++
		} else if r == len(arr) {
			l--
		} else {
			if abs(arr[l], x) <= abs(arr[r], x) {
				if l != -1 {
					l--
				} else {
					r++
				}

			} else {
				if r != len(arr) {
					r++
				} else {
					l--
				}
			}
		}
	}
	return arr[l+1 : r]
}
func abs(a, b int) int {
	if a >= b {
		return a - b
	} else {
		return b - a
	}
}

// @lc code=end
func main() {

	fmt.Println(findClosestElements([]int{3, 5, 8, 10}, 2, 15))
	fmt.Println(findClosestElements([]int{0, 0, 1, 2, 3, 3, 4, 7, 7, 8}, 3, 5))
	fmt.Println(findClosestElements([]int{1, 2, 3, 4, 5}, 4, 3))
	fmt.Println(findClosestElements([]int{1, 2, 3, 4, 5}, 4, -1))
}
