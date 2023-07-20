package main

import (
	"fmt"
	"strconv"
	"time"
)
// 采用sleep形式防止主协程执行结束，导致三个协程退出
// 定义 3 个函数分别打印 cat、dog、fish，要求每个函数都要起一个 goroutine，按照 cat、dog、fish 顺序打印在屏幕上 100 次。
func main() {
	catChan := make(chan struct{}, 1)
	dogChan := make(chan struct{}, 1)
	fishChan := make(chan struct{}, 1)
	catChan <- struct{}{}

	go printCat(catChan, dogChan)
	go printDog(dogChan, fishChan)
	go printFish(fishChan, catChan)

	time.Sleep(time.Second)
}
func printCat(catChan, dogChan chan struct{}) {
	for i := 0; i < 100; i++ {
		<-catChan
		fmt.Println("cat" + strconv.Itoa(i))
		dogChan <- struct{}{}
	}
}
func printDog(dogChan, fishChan chan struct{}) {
	for i := 0; i < 100; i++ {
		<-dogChan
		fmt.Println("dog" + strconv.Itoa(i))
		fishChan <- struct{}{}
	}
}
func printFish(fishChan, catChan chan struct{}) {
	for i := 0; i < 100; i++ {
		<-fishChan
		fmt.Println("fish" + strconv.Itoa(i))
		catChan <- struct{}{}
	}
}
