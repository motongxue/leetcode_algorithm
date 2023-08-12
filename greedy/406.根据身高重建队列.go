/*
 * @lc app=leetcode.cn id=406 lang=golang
 *
 * [406] 根据身高重建队列
 */
// TODO 核心是先按身高进行倒序排序，身高相同，则按k进行正序排序
// 然后，插入位置按h进行插入
// @lc code=start
func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})
	for i, p := range people {
		copy(people[p[1]+1:i+1], people[p[1]:i+1])
		people[p[1]] = p
	}
	// 同这一步
	// for (int[] p : people) {
	// 	que.add(p[1],p);
	// }
	return people
}

// ================================================================================
// 重写
// 采用插入排序的思想，注意，这里的插入排序是指将元素按照ki进行插入到指定位置，而不是在原数组位置上进行排序
func reconstructQueue(people [][]int) [][]int {
	// 按身高hi进行降序排序，并且需要对ki进行升序排序
	sort.Slice(people, func(a, b int) bool {
		if people[a][0] == people[b][0] {
			return people[a][1] < people[b][1]
		}
		return people[a][0] > people[b][0]
	})
	// 初始化res返回结果
	n := len(people)
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, 2)
	}
	// 插入排序
	for i := 0; i < n; i++ {
		insert(people[i], res)
	}
	return res
}
func insert(v []int, res [][]int) {
	n := len(res)
	i := n - 1
	for i = n - 1; i > v[1]; i-- {
		res[i] = res[i-1]
	}
	res[i] = v
}

// @lc code=end

