/*
 * @lc app=leetcode.cn id=59 lang=golang
 *
 * [59] 螺旋矩阵 II
 */
package main

import "fmt"

// @lc code=start
// TODO 注意到了一个方向的终点才换方向
func generateMatrix(n int) [][]int {
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
	result := [20][20]int{}
	result[sx][sy] = 1
	n2 := n * n
	for cntI < n2 {
		sx = sx + dir[dirI][0]
		sy = sy + dir[dirI][1]
		// 到达，更换方向
		if sy < 0 || sx < 0 || sy == n || sx == n || vis[sx][sy] {
			sx = sx - dir[dirI][0]
			sy = sy - dir[dirI][1]
			dirI = (dirI + 1) % 4
			continue
		}
		vis[sx][sy] = true
		cntI++
		result[sx][sy] = cntI
	}
	resultSlice := [][]int{}
	for i := 0; i < n; i++ {
		tmp := []int{}
		for j := 0; j < n; j++ {
			tmp = append(tmp, result[i][j])
		}
		resultSlice = append(resultSlice, tmp)
	}
	return resultSlice
}

// @lc code=end
func main() {
	fmt.Println(generateMatrix(3))
}
