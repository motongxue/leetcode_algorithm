/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */
package main

// @lc code=start
// 采用单调队列来实现
// 队列没有必要维护窗口里的所有元素，
// 只需要维护有可能成为窗口里最大值的元素就可以了，
// 同时保证队列里的元素数值是由大到小的。
type MonotonicQueue struct {
	q []int
}

func (mq *MonotonicQueue) push(n int) {
	for len(mq.q) > 0 && mq.q[len(mq.q)-1] < n {
		mq.q = mq.q[:len(mq.q)-1]
	}
	mq.q = append(mq.q, n)
}
func (mq *MonotonicQueue) pop(n int) {
	if n == mq.q[0] {
		mq.q = mq.q[1:]
	}
}
func maxSlidingWindow(nums []int, k int) []int {
	window := &MonotonicQueue{}
	res := []int{}
	for i := 0; i < k; i++ {
		window.push(nums[i])
	}
	res = append(res, window.q[0])
	for i := 0; i+k < len(nums); i++ {
		window.pop(nums[i])
		window.push(nums[i+k])
		res = append(res, window.q[0])
	}
	return res
}

// @lc code=end
