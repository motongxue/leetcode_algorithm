/*
 * @lc app=leetcode.cn id=225 lang=golang
 *
 * [225] 用队列实现栈
 */

// @lc code=start
// // 采用两个队列实现栈
// type MyStack struct {
// 	s1 []int
// 	s2 []int
// }

// func Constructor() MyStack {
// 	return MyStack{}
// }

// func (this *MyStack) Push(x int) {
// 	if len(this.s2) != 0 {
// 		this.s1 = this.s2
// 		this.s2 = []int{}
// 	}
// 	this.s1 = append(this.s1, x)
// }

// func (this *MyStack) Pop() int {
// 	for len(this.s1) != 1 {
// 		this.s2 = append(this.s2, this.s1[0])
// 		this.s1 = this.s1[1:len(this.s1)]
// 	}
// 	t := this.s1[0]
// 	this.s1 = this.s2
// 	this.s2 = []int{}
// 	return t
// }

// func (this *MyStack) Top() int {
// 	for len(this.s1) != 1 {
// 		this.s2 = append(this.s2, this.s1[0])
// 		this.s1 = this.s1[1:len(this.s1)]
// 	}
// 	t := this.s1[0]
// 	this.s2 = append(this.s2, t)
// 	this.s1 = this.s2
// 	this.s2 = []int{}
// 	return t
// }

// func (this *MyStack) Empty() bool {
// 	if len(this.s1) == 0 && len(this.s2) == 0 {
// 		return true
// 	}
// 	return false
// }

// 采用两个队列实现栈
type MyStack struct {
	queue []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.queue = append(this.queue, x)
}

func (this *MyStack) Pop() int {
	n := len(this.queue)
	for i := 0; i < n-1; i++ {
		tmp := this.queue[0]
		this.queue = this.queue[1:n]
		this.queue = append(this.queue, tmp)
	}
	tmp := this.queue[0]
	this.queue = this.queue[1:len(this.queue)]
	return tmp
}

func (this *MyStack) Top() int {
	n := len(this.queue)
	for i := 0; i < n-1; i++ {
		tmp := this.queue[0]
		this.queue = this.queue[1:n]
		this.queue = append(this.queue, tmp)
	}
	tmp := this.queue[0]
	this.queue = this.queue[1:len(this.queue)]
	this.queue = append(this.queue, tmp)
	return tmp
}

func (this *MyStack) Empty() bool {
	if len(this.queue) == 0 {
		return true
	}
	return false
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
// @lc code=end

