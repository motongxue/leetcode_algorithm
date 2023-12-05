// 完全背包问题，每个物品可以无限次选择
// ================================================================
// 先遍历物品，再遍历背包
func UnboundedKnapsackProblem_1() {
	weight := []int{1, 3, 4}   // 物品重量
	value := []int{15, 20, 30} // 物品价值
	bagWeight := 4             // 背包容量
	// 初始化dp数组
	dp := make([]int, bagWeight+1)
	for i := 0; i < len(weight); i++ {
		for j := weight[i]; j <= bagWeight; j++ {
			dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
		}
	}
	fmt.Printf("完全背包问题，每个物品可以无限次选择，最终结果为：%d\n", dp[bagWeight])
}

// ================================================================
// 先遍历背包，再遍历物品
func UnboundedKnapsackProblem_2() {
	weight := []int{1, 3, 4}   // 物品重量
	value := []int{15, 20, 30} // 物品价值
	bagWeight := 4             // 背包容量
	// 初始化dp数组
	dp := make([]int, bagWeight+1)
	for j := 0; j <= bagWeight; j++ {
		for i := 0; i < len(weight); i++ {
			if j >= weight[i] {
				dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
			}
		}
	}
	fmt.Printf("完全背包问题，每个物品可以无限次选择，最终结果为：%d\n", dp[bagWeight])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
