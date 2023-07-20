package main

import (
	"fmt"
	"strconv"
	"sync"
)
// 采用匿名函数形式
// 定义 3 个函数分别打印 cat、dog、fish，要求每个函数都要起一个 goroutine，按照 cat、dog、fish 顺序打印在屏幕上 100 次。
func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	catChan := make(chan struct{}, 1)
	dogChan := make(chan struct{}, 1)
	fishChan := make(chan struct{}, 1)
	go func() {
		for i := 0; i < 100; i++ {
			<-catChan
			fmt.Println("cat" + strconv.Itoa(i))
			dogChan <- struct{}{}
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 100; i++ {
			<-dogChan
			fmt.Println("dog" + strconv.Itoa(i))
			fishChan <- struct{}{}
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 100; i++ {
			// wg.Add(1)
			<-fishChan
			fmt.Println("fish" + strconv.Itoa(i))
			catChan <- struct{}{}
		}
		wg.Done()
	}()
	catChan <- struct{}{}
	wg.Wait()
}
