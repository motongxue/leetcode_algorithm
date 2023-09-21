// 01背包问题，每个物品只能选一次
// ==================
// 二维数组情况：由于是根据左边和上边的值来计算当前值，所以需要先遍历物品或者先遍历背包容量都可以
func zero_one_package_2() {
	weight := []int{1, 3, 4}   // 物品重量
	value := []int{15, 20, 30} // 物品价值
	bagWeight := 4             // 背包容量
	// 初始化二维dp数组
	dp := make([][]int, len(weight))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, bagWeight+1)
	}
	// 对边界赋值
	for j := weight[0]; j <= bagWeight; j++ {
		dp[0][j] = value[0]
	}
	for i := 1; i < len(weight); i++ {
		for j := 0; j <= bagWeight; j++ {
			// 如果当前背包容量j大于等于物品i的重量，那么可以选择装或者不装
			if j >= weight[i] {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	fmt.Println("01背包问题，每个物品只能选一次，最终结果为：%d\n", dp[len(weight)-1][bagWeight])
}

// ==================
// 一维数组情况：倒序遍历是为了保证物品i只被放入一次，并且保证实现滚动数组，防止dp[i-1][j-weight[i]]被覆盖
func zero_one_package_1() {
	weight := []int{1, 3, 4}   // 物品重量
	value := []int{15, 20, 30} // 物品价值
	bagWeight := 4             // 背包容量
	dp := make([]int, bagWeight+1)
	for i := 0; i < len(weight); i++ {
		// 倒序遍历
		for j := bagWeight; j >= weight[i]; j-- {
			dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
		}
	}
	fmt.Printf("01背包问题，每个物品只能选一次，最终结果为：%d\n", dp[bagWeight])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}