// 1114. 按序打印

// 三个不同的线程 A、B、C 将会共用一个 Foo 实例。
// 线程 A 将会调用 first() 方法
// 线程 B 将会调用 second() 方法
// 线程 C 将会调用 third() 方法

// 输入：nums = [1,2,3]
// 输出："firstsecondthird"
// 输入：nums = [1,3,2]
// 输出："firstsecondthird"

package main

import (
	"fmt"
)

// 这种方式属于执行一次函数，便销毁一次函数，所以不需要等待信号量
func first(ch chan struct{}) {
	fmt.Println("1")
	ch <- struct{}{}
}

func second(ch chan struct{}) {
	fmt.Println("2")
	ch <- struct{}{}
}

func third(ch chan struct{}) {
	fmt.Println("3")
	ch <- struct{}{}
}

func main() {
	nums := []int{2, 1, 3, 2, 1, 3, 1, 2, 3}

	ch := make(chan struct{}, 1)
	ch <- struct{}{}

	for _, num := range nums {
		<-ch
		switch num {
		case 1:
			go first(ch)
		case 2:
			go second(ch)
		case 3:
			go third(ch)
		}
	}
	// 防止主协程执行结束
	<-ch
	// 释放channel
	close(ch)
}
