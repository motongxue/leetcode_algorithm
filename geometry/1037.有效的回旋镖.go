/*
 * @lc app=leetcode.cn id=1037 lang=golang
 *
 * [1037] 有效的回旋镖
 */
package main

import "fmt"

// @lc code=start
func isBoomerang(points [][]int) bool {
    // y2-y1/x2-x1=y3-y1/x3-x1，但存在为零情况，故相乘
    v1 := [2]int{points[1][0] - points[0][0], points[1][1] - points[0][1]}
    v2 := [2]int{points[2][0] - points[0][0], points[2][1] - points[0][1]}
    return v1[0]*v2[1]-v1[1]*v2[0] != 0
}
// @lc code=end
func main() {
	fmt.Println(isBoomerang([][]int{{0, 1}, {0, 1}, {2, 1}}))
}
