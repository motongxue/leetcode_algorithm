/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 */

// @lc code=start
func jump(nums []int) int {
	res := 0
	curDistance := 0
	nextDistance := 0
	if len(nums) == 1 {
		return 0
	}
	for i := 0; i < len(nums); i++ {
		nextDistance = max(nums[i]+i, nextDistance)
		if nextDistance >= len(nums)-1 {
			res++
			break
		}
		if i == curDistance {
			curDistance = nextDistance
			res++
		}
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

