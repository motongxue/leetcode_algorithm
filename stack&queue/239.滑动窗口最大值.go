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
// 定义一：
type MonotonicQueue struct {
	q []int
}

func (mq *MonotonicQueue) push(n int) {
	// 当切片长度不为零，且n大于切片数组最后一位时，则直接将最后一位删除，使得切片中第0位是最大的
	for len(mq.q) > 0 && mq.q[len(mq.q)-1] < n {
		mq.q = mq.q[:len(mq.q)-1]
	}
	mq.q = append(mq.q, n)
}
/* // 定义二：
type MQ []int
func (mq *MQ) push(in int) {
    for len(*mq) != 0 && (*mq)[len(*mq)-1] < in {
        *mq = (*mq)[:len(*mq)-1]
    }
    *mq = append(*mq, in)
}
func (mq *MQ) pop(in int) {
    if in == (*mq)[0] {
        *mq = (*mq)[1:]
    }
}
 */

// 当弹出最大的元素时，才对队列中的元素进行修改
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
