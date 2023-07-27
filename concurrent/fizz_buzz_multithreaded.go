// 1195. 交替打印字符串
// 编写一个可以从 1 到 n 输出代表这个数字的字符串的程序，但是：

// 如果这个数字可以被 3 整除，输出 "fizz"。
// 如果这个数字可以被 5 整除，输出 "buzz"。
// 如果这个数字可以同时被 3 和 5 整除，输出 "fizzbuzz"。
// 例如，当 n = 15，输出： 1, 2, fizz, 4, buzz, fizz, 7, 8, fizz, buzz, 11, fizz, 13, 14, fizzbuzz。

// 假设有这么一个类：
// class FizzBuzz {
//   public FizzBuzz(int n) { ... }               // constructor
//   public void fizz(printFizz) { ... }          // only output "fizz"
//   public void buzz(printBuzz) { ... }          // only output "buzz"
//   public void fizzbuzz(printFizzBuzz) { ... }  // only output "fizzbuzz"
//   public void number(printNumber) { ... }      // only output the numbers
// }
// 请你实现一个有四个线程的多线程版  FizzBuzz， 同一个 FizzBuzz 实例会被如下四个线程使用：

// 线程A将调用 fizz() 来判断是否能被 3 整除，如果可以，则输出 fizz。
// 线程B将调用 buzz() 来判断是否能被 5 整除，如果可以，则输出 buzz。
// 线程C将调用 fizzbuzz() 来判断是否同时能被 3 和 5 整除，如果可以，则输出 fizzbuzz。
// 线程D将调用 number() 来实现输出既不能被 3 整除也不能被 5 整除的数字。
package main

import (
	"fmt"
	"sync"
	"time"
)

// 定义结构体
type FizzBuzz struct {
	n int
	// 为了控制并发
	fizzChan     chan struct{}
	buzzChan     chan struct{}
	fizzbuzzChan chan struct{}
	numberChan   chan struct{}
	wg           sync.WaitGroup
}
var res string
func NewFizzBuzz(n int) *FizzBuzz {
	return &FizzBuzz{
		n:            n,
		fizzChan:     make(chan struct{}, 1),
		buzzChan:     make(chan struct{}, 1),
		fizzbuzzChan: make(chan struct{}, 1),
		numberChan:   make(chan struct{}, 1),
		wg:           sync.WaitGroup{},
	}
}

func (f *FizzBuzz) fizz() {
	for {
		<-f.fizzChan
		res += "fizz\t"
		fmt.Printf("fizz\t")
		f.wg.Done()
	}
}

func (f *FizzBuzz) buzz() {
	for {
		<-f.buzzChan
		res += "buzz\t"
		fmt.Printf("buzz\t")
		f.wg.Done()
	}
}

func (f *FizzBuzz) fizzbuzz() {
	for {
		<-f.fizzbuzzChan
		res += "fizzbuzz\t"
		fmt.Printf("fizzbuzz\t")
		f.wg.Done()
	}
}

func (f *FizzBuzz) number() {
	for {
		<-f.numberChan
		res += fmt.Sprintf("%d\t", f.n)
		fmt.Printf("%d\t", f.n)
		f.wg.Done()
	}
}

func main() {
	n := 15
	f := NewFizzBuzz(n)
	f.wg.Add(n)
	defer f.wg.Wait()
	go f.fizz()
	go f.buzz()
	go f.fizzbuzz()
	go f.number()
	for i := 1; i <= n; i++ {
		f.n = i
		if i%5 == 0 && i%3 == 0 {
			f.fizzbuzzChan <- struct{}{}
		} else if i%3 == 0 {
			f.fizzChan <- struct{}{}
		} else if i%5 == 0 {
			f.buzzChan <- struct{}{}
		} else {
			f.numberChan <- struct{}{}
		}
		// 使得输出有序
		time.Sleep(time.Millisecond)
	}
}
