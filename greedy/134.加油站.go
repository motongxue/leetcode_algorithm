/*
 * @lc app=leetcode.cn id=134 lang=golang
 *
 * [134] 加油站
 */
// TODO　不理解
// 核心点：亏空最严重的一个点必须放在最后一步走，等着前面剩余的救助
// @lc code=start
func canCompleteCircuit(gas []int, cost []int) int {
    spare, minSpare, minSpareIdx := 0, math.MaxInt, 0
    for i := 0; i < len(gas); i++ {
        spare += gas[i] - cost[i]
        if spare < minSpare  {
            minSpare = spare
            minSpareIdx = i
        }
    }
    if spare < 0 {
        return -1
    }
    return (minSpareIdx + 1) % len(gas)
}
// @lc code=end

