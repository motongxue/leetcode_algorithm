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
// 如果是这种方式，则上限为cover
// func canJump(nums []int) bool {
//     cover := 0
//     n := len(nums)-1
//     for i := 0; i <= cover; i++ { // 每次与覆盖值比较
//         fmt.Printf("%d %d \n",i+nums[i],cover)
//         cover = max(i+nums[i], cover) //每走一步都将 cover 更新为最大值
//         if cover >= n {
//             return true
//         }
//     }
//     return false
// }
// @lc code=end

