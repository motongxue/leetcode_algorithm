/*
 * @lc app=leetcode.cn id=376 lang=golang
 *
 * [376] 摆动序列
 */

// @lc code=start
// ========================================
// 贪心算法实现
// func wiggleMaxLength(nums []int) int {
//     if len(nums) == 1 {
//         return len(nums)
//     }
//     res := 1
//     preDiff, curDiff := 0, 0
//     for i := 1; i < len(nums); i++ {
//         curDiff = nums[i] - nums[i-1]
//         if preDiff <= 0 && curDiff > 0 || preDiff >= 0 && curDiff < 0 {
//             res++
//             preDiff = curDiff; // 注意这里，只在摆动变化的时候更新prediff 
//         }
//         // 在平坡的时候不需要改动: 1 2 2 2 3 4 实际上只有两个峰值
// 		// 所以需要把更新操作放到上面
//         // preDiff = curDiff
//     }
//     return res
// }
// 自己实现
func wiggleMaxLength(nums []int) int {
    n := len(nums)
    if n <= 1 {
        return n
    }
    preDiff, curDiff := 0, 0
    cnt := 1
    // 就是将preDiff和curDiff的各种情况列出来，然后判断哪种情况才需要更新
    for i := 1; i < n; i++ {
        curDiff = nums[i]-nums[i-1]
        if (preDiff > 0 && curDiff < 0) || (preDiff < 0 && curDiff > 0) || (preDiff == 0 && curDiff != 0) {
            preDiff = curDiff
            cnt++
        }
    }
    return cnt
}

// TODO 动态规划实现
// TODO 线段树实现
// @lc code=end

