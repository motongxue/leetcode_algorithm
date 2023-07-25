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
// ===================================================================
// 新增更简洁的解法，并添加注解
func jump(nums []int) int {
    // nextStep: 目前能跳到的最远位置，curStep: 上次跳跃可达范围右边界
    nextStep, curStep := 0, 0
    res, n := 0, len(nums)
    // 到达n-1即为到达，故可以只考虑i<n-1
    for i := 0; i < n-1; i++ {
        // 不断更新目前能跳到的最远距离
        nextStep = max(nextStep, nums[i]+i)
        // 如果此时已经到达上次的最右边界，则更新
        if i == curStep {
            res++
            curStep = nextStep
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

