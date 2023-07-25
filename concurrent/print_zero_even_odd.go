// 1116. 打印零与奇偶数
// 现有函数 printNumber 可以用一个整数参数调用，并输出该整数到控制台。
// 给你类 ZeroEvenOdd 的一个实例，该类中有三个函数：zero、even 和 odd 。ZeroEvenOdd 的相同实例将会传递给三个不同线程：

// 线程 A：调用 zero() ，只输出 0
// 线程 B：调用 even() ，只输出偶数
// 线程 C：调用 odd() ，只输出奇数
// 修改给出的类，以输出序列 "010203040506..." ，其中序列的长度必须为 2n 。

// 实现 ZeroEvenOdd 类：

// ZeroEvenOdd(int n) 用数字 n 初始化对象，表示需要输出的数。
// void zero(printNumber) 调用 printNumber 以输出一个 0 。
// void even(printNumber) 调用printNumber 以输出偶数。
// void odd(printNumber) 调用 printNumber 以输出奇数。
// 输入：n = 5
// 输出："0102030405"
package main

import (
	"fmt"
	"sync"
)

func main() {
	zeroChan := make(chan struct{}, 1)
	evenChan := make(chan struct{}, 1)
	oddChan := make(chan struct{}, 1)
	var n int
	fmt.Scanln(&n)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go zero(n, zeroChan, evenChan, oddChan, &wg)
	go even(n, evenChan, zeroChan)
	go odd(n, oddChan, zeroChan)
	zeroChan <- struct{}{}
	wg.Wait()
	// 防止主线程提前退出
	<-zeroChan
}

func zero(n int, zeroChan, evenChan, oddChan chan struct{}, wg *sync.WaitGroup) {
	for i := 0; i < n; i++ {
		<-zeroChan
		fmt.Print(0)
		if i&1 == 0 {
			oddChan <- struct{}{}
		} else {
			evenChan <- struct{}{}
		}
	}
	defer wg.Done()
}

func even(n int, evenChan, zeroChan chan struct{}) {
	// 由于是偶数且为2*i，所以从1开始
	for i := 1; i <= n; i++ {
		<-evenChan
		fmt.Print(2 * i)
		zeroChan <- struct{}{}
	}
}

func odd(n int, oddChan, zeroChan chan struct{}) {
	// 由于是奇数且为2*i+1，所以从0开始
	for i := 0; i < n; i++ {
		<-oddChan
		fmt.Print(2*i + 1)
		zeroChan <- struct{}{}
	}
}
