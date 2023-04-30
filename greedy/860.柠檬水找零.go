/*
 * @lc app=leetcode.cn id=860 lang=golang
 *
 * [860] 柠檬水找零
 */

// @lc code=start
func lemonadeChange(bills []int) bool {
    mp := make(map[int]int)
    for _, v := range bills {
        if v == 5 {
            mp[5]++
        }else if v == 10 {
            if mp[5] > 0 {
                mp[5]--
                mp[10]++
            }else {
                return false
            }
        }else {
			// 为什么要优先消耗一个10和一个5呢？
			// 因为美元10只能给账单20找零，而美元5可以给账单10和账单20找零，美元5更万能！
            if mp[5] >= 1 && mp[10] >= 1 {
                mp[5]--
                mp[10]--
                mp[20]++
            }else if mp[5] >= 3 {
                mp[5]-=3
                mp[20]++
            } else {
                return false
            }
        }
    }
    return true
}
// @lc code=end

