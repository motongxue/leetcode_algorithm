/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 */
package main

import "fmt"

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	n := len(matrix)
	m := len(matrix[0])
	dir := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	dirI := 0
	sx, sy := 0, 0
	cntI := 1

	vis := [20][20]bool{}
	vis[sx][sy] = true
	result := []int{}
	result = append(result, matrix[0][0])
	n2 := n * m
	for cntI < n2 {
		sx = sx + dir[dirI][0]
		sy = sy + dir[dirI][1]
		// 到达，更换方向
		if sy < 0 || sx < 0 || sy == m || sx == n || vis[sx][sy] {
			sx = sx - dir[dirI][0]
			sy = sy - dir[dirI][1]
			dirI = (dirI + 1) % 4
			continue
		}
		vis[sx][sy] = true
		cntI++
		result = append(result, matrix[sx][sy])
	}
	return result
}

// @lc code=end

func main() {
	fmt.Println(spiralOrder([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}))
	fmt.Println(spiralOrder([][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}))
}
