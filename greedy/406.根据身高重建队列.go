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
// @lc code=end

