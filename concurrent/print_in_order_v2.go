// 1114. 按序打印

// 三个不同的线程 A、B、C 将会共用一个 Foo 实例。
// 线程 A 将会调用 first() 方法
// 线程 B 将会调用 second() 方法
// 线程 C 将会调用 third() 方法

// 输入：nums = [1,2,3]
// 输出："firstsecondthird"
// 输入：nums = [1,3,2]
// 输出："firstsecondthird"
// 请为我修复下面这段代码，使得输入 nums 所代表的函数调用能够正确地输出 "firstsecondthird"。

// ===================================================================
// 实际需要的效果：函数只需启动一次，利用 for + 监听信号量的形式实现按序打印
package main

import (
	"fmt"
)

var (
	// 定义信号量
	firstChan  = make(chan struct{}, 1)
	secondChan = make(chan struct{}, 1)
	thirdChan  = make(chan struct{}, 1)
	ch         = make(chan struct{}, 1)
)

func first() {
	for {
		<-firstChan
		fmt.Println("1")
		ch <- struct{}{}
	}
}

func second() {
	for {
		<-secondChan
		fmt.Println("2")
		ch <- struct{}{}
	}
}

func third() {
	for {
		<-thirdChan
		fmt.Println("3")
		ch <- struct{}{}
	}
}

func main() {
	nums := []int{2, 1, 3, 2, 1, 3, 1, 2, 3}
	go first()
	go second()
	go third()
	ch <- struct{}{}
	for _, num := range nums {
		<-ch
		switch num {
		case 1:
			firstChan <- struct{}{}
		case 2:
			secondChan <- struct{}{}
		case 3:
			thirdChan <- struct{}{}
		}
	}
	// 防止主协程执行结束
	<-ch
	// 释放channel
	close(ch)
	close(firstChan)
	close(secondChan)
	close(thirdChan)
}
