/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */

// @lc code=start
func canJump(nums []int) bool {
	step := 0
	for i, v := range nums {
		// i = step && v = 0 表明它最远只能跳到这里
		// 但比如 [0]，还要判断他是否已经到达末尾
		if i == step && v == 0 && i != len(nums)-1 {
			return false
		}
		step = max(step, i+v)
	}
	return true
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
// ============================================
// 记录数组中每个位置所能达到的最大值，然后i<maxStep，最后判断maxStep是否达到了数组的最后一个元素位置
func canJump(nums []int) bool {
    maxStep := nums[0]
    for i := 0; i <= maxStep && i < len(nums); i++ {
        maxStep = max(i + nums[i], maxStep)
    }
    return maxStep >= len(nums)-1
}
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
// @lc code=end

