/*
 * @lc app=leetcode.cn id=232 lang=golang
 *
 * [232] 用栈实现队列
 */
package main

// @lc code=start
// 用两个栈来实现队列
// push的时候，入s2，出栈的时候，若s1为空，需要将s2中所有元素添加到s1中，然后从s1中返回
type MyQueue struct {
	stackOut, stackIn []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.stackIn = append(this.stackIn, x)
}

func (this *MyQueue) Pop() int {
	var t int
	if len(this.stackOut) != 0 {
		t = this.stackOut[len(this.stackOut)-1]
		this.stackOut = this.stackOut[:len(this.stackOut)-1]
	} else {
		for i:=len(this.stackIn)-1;i>=0;i--{
			this.stackOut = append(this.stackOut, this.stackIn[i])
		}
		this.stackIn = []int{}
		t = this.stackOut[len(this.stackOut)-1]
		this.stackOut = this.stackOut[:len(this.stackOut)-1]
	}
	return t
}

func (this *MyQueue) Peek() int {
	var t int
	if len(this.stackOut) != 0 {
		t = this.stackOut[len(this.stackOut)-1]
	} else {
		for i:=len(this.stackIn)-1;i>=0;i--{
			this.stackOut = append(this.stackOut, this.stackIn[i])
		}
		this.stackIn = []int{}
		t = this.stackOut[len(this.stackOut)-1]
	}
	return t
}

func (this *MyQueue) Empty() bool {
	if len(this.stackOut) == 0 && len(this.stackIn) == 0 {
		return true
	}
	return false
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
// @lc code=end
