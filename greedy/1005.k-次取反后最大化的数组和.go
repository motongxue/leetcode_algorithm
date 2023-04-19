/*
 * @lc app=leetcode.cn id=1005 lang=golang
 *
 * [1005] K 次取反后最大化的数组和
 */
//  TODO 重做
// @lc code=start
func largestSumAfterKNegations(nums []int, k int) int {
	// 先通过绝对值来排序，由大到小
	sort.Slice(nums, func(a, b int) bool {
		return abs(nums[a]) > abs(nums[b])
	})
	for i := 0; i < len(nums); i++ {
		// 从大的开始，如果是负数，且还有k，则该数取负
		if nums[i] < 0 && k > 0 {
			nums[i] *= -1
			k--
		}
	}
	// 如果k没剩，那说明能转的负数都转正了，已经是最大和，返回sum
	// 如果k有剩，说明负数已经全部转正，所以如果k还剩偶数个就自己抵消掉，不用删减，如果k还剩奇数个就减掉2倍最小正数。
	if k%2 == 1 {
		nums[len(nums)-1] *= -1
	}
	res := 0
	for _, v := range nums {
		res += v
	}
	return res
}
func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

// @lc code=end

